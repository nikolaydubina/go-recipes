package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/nikolaydubina/mdpage/page"
	"github.com/nikolaydubina/mdpage/render"
	yaml "gopkg.in/yaml.v3"
)

func loadPageYAML(t *testing.T, path string) page.Page {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}

	var p page.Page
	if err := yaml.NewDecoder(bytes.NewReader(data)).Decode(&p); err != nil {
		t.Fatalf("failed to parse yaml in %s: %v", path, err)
	}
	return p
}

func TestPageYAMLValidation(t *testing.T) {
	p := loadPageYAML(t, "page.yaml")

	if err := p.Validate(); err != nil {
		t.Fatalf("page.yaml validation failed: %v", err)
	}

	if len(p.Groups) == 0 {
		t.Fatal("expected at least one group in page.yaml")
	}

	var totalEntries int
	for _, g := range p.Groups {
		if g.Title == "" {
			t.Errorf("group has empty title: %+v", g)
		}
		if len(g.Entries) == 0 {
			t.Errorf("group %q has no entries", g.Title)
		}
		totalEntries += len(g.Entries)
	}

	if totalEntries == 0 {
		t.Fatal("expected entries in page.yaml, got 0")
	}
}

func TestShapeOfGoEntry_PresentAndValid(t *testing.T) {
	p := loadPageYAML(t, "page.yaml")

	var foundEntry *page.Entry
	var foundInGroup string

	for _, g := range p.Groups {
		for _, e := range g.Entries {
			if e.Name == "shape-of-go" {
				copyEntry := e
				foundEntry = &copyEntry
				foundInGroup = g.Title
				break
			}
		}
	}

	if foundEntry == nil {
		t.Fatal("shape-of-go entry not found in page.yaml")
	}

	if foundInGroup != "Dependencies" {
		t.Errorf("expected shape-of-go to be in Dependencies group, got %q", foundInGroup)
	}

	if expected := "https://github.com/LuisLSousa/shape-of-go"; foundEntry.URL != expected {
		t.Errorf("expected url %q, got %q", expected, foundEntry.URL)
	}

	if expected := "https://github.com/LuisLSousa"; foundEntry.Author != expected {
		t.Errorf("expected author %q, got %q", expected, foundEntry.Author)
	}

	if expected := "Explore Go module ecosystem dependency graph interactively"; foundEntry.Title != expected {
		t.Errorf("expected title %q, got %q", expected, foundEntry.Title)
	}

	if expected := "https://raw.githubusercontent.com/LuisLSousa/shape-of-go/main/docs/images/galaxy-overview.jpg"; foundEntry.ExampleImageURL != expected {
		t.Errorf("expected example_image_url %q, got %q", expected, foundEntry.ExampleImageURL)
	}

	if !strings.Contains(foundEntry.Description, "shape-of-go.pages.dev") {
		t.Errorf("expected description to mention shape-of-go.pages.dev: %q", foundEntry.Description)
	}

	if !strings.Contains(foundEntry.Description, "2.6 million modules") {
		t.Errorf("expected description to mention module statistics: %q", foundEntry.Description)
	}
}

func TestREADMEContainsShapeOfGo(t *testing.T) {
	content, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("failed to read README.md: %v", err)
	}
	text := string(content)

	expectedTOC := "[ Explore Go module ecosystem dependency graph interactively with `shape-of-go`](#-explore-go-module-ecosystem-dependency-graph-interactively-with-shape-of-go)"
	if !strings.Contains(text, expectedTOC) {
		t.Errorf("README.md missing table of contents link: %q", expectedTOC)
	}

	expectedHeader := "### [⏫](#contents) Explore Go module ecosystem dependency graph interactively with [shape-of-go](https://github.com/LuisLSousa/shape-of-go)"
	if !strings.Contains(text, expectedHeader) {
		t.Errorf("README.md missing section header: %q", expectedHeader)
	}

	expectedImage := `<img src="https://raw.githubusercontent.com/LuisLSousa/shape-of-go/main/docs/images/galaxy-overview.jpg"`
	if !strings.Contains(text, expectedImage) {
		t.Errorf("README.md missing expected image: %q", expectedImage)
	}

	expectedAuthor := "[@LuisLSousa](https://github.com/LuisLSousa)"
	if !strings.Contains(text, expectedAuthor) {
		t.Errorf("README.md missing author link: %q", expectedAuthor)
	}

	expectedSite := "[shape-of-go.pages.dev](https://shape-of-go.pages.dev/)"
	if !strings.Contains(text, expectedSite) {
		t.Errorf("README.md missing live site link: %q", expectedSite)
	}
}

func TestRenderedOutputMatchesREADME(t *testing.T) {
	p := loadPageYAML(t, "page.yaml")
	if err := p.Validate(); err != nil {
		t.Fatalf("page validation failed: %v", err)
	}

	renderer := render.NewSimplePageRender(render.GitHubAuthorRenderer{Prefix: "@"})
	var buf strings.Builder
	renderer.RenderTo(&buf, p)

	readmeBytes, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("failed to read README.md: %v", err)
	}

	if buf.String() != string(readmeBytes) {
		t.Error("README.md is out of sync with page.yaml rendering")
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("non-existent file", func(t *testing.T) {
		_, err := os.ReadFile("non_existent_page.yaml")
		if err == nil {
			t.Fatal("expected error reading non-existent file, got nil")
		}
	})

	t.Run("invalid yaml syntax", func(t *testing.T) {
		invalidYAML := []byte("title: [unclosed list")
		var p page.Page
		err := yaml.NewDecoder(bytes.NewReader(invalidYAML)).Decode(&p)
		if err == nil {
			t.Fatal("expected error decoding invalid yaml, got nil")
		}
	})

	t.Run("empty entry validation", func(t *testing.T) {
		var p page.Page
		p.Groups = []page.Group{
			{
				Title: "Test Group",
				Entries: []page.Entry{
					{},
				},
			},
		}
		if err := p.Validate(); err == nil {
			t.Error("expected validation error for empty entry, got nil")
		}
	})
}
