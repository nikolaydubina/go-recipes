<h1 align="center">Go Recipes 🦩 </h1>
<p align="center">Handy well-known and <i>lesser</i>-known tools for Go projects</p>

> _Know some cool tool or one-liner? Have a feature request or an idea?_  
> _Feel free to edit this page or create an Issue/Discussion!_  

## Contents

 - Tests
   + [➡ Make treemap of code coverage](#-make-treemap-of-code-coverage)
   + [➡ Pretty print coverage of Go code in terminal](#-pretty-print-coverage-of-go-code-in-terminal)
   + [➡ Browse code coverage of Go code in terminal](#-browse-code-coverage-of-go-code-in-terminal)
   + [➡ Get packages without tests](#-get-packages-without-tests)
   + [➡ Browse code coverage by file](#-browse-code-coverage-by-file)
   + [➡ Make histogram of Go files per package](#-make-histogram-of-go-files-per-package)
   + [➡ Run tests sequentially](#-run-tests-sequentially)
   + [➡ Run tests in parallel](#-run-tests-in-parallel)
   + [➡ Run tests with pretty output](#-run-tests-with-pretty-output)
   + [➡ Detect goroutine leaks](#-detect-goroutine-leaks)
 - Dependencies
   + [➡ Get Go version of current module](#-get-go-version-of-current-module)
   + [➡ Get Go versions of upstream modules](#-get-go-versions-of-upstream-modules)
   + [➡ Get directly dependent modules that can be upgraded](#-get-directly-dependent-modules-that-can-be-upgraded)
   + [➡ Get upstream modules without Go version](#-get-upstream-modules-without-go-version)
   + [➡ Get available module versions](#-get-available-module-versions)
   + [➡ Make graph of upstream modules](#-make-graph-of-upstream-modules)
   + [➡ Make graph of upstream modules with gmchart](#-make-graph-of-upstream-modules-with-gmchart)
   + [➡ Make graph of upstream packages](#-make-graph-of-upstream-packages)
   + [➡ Scrape details about upstream modules and make graph](#-scrape-details-about-upstream-modules-and-make-graph)
   + [➡ Scrape licenses of upstream dependencies](#-scrape-licenses-of-upstream-dependencies)
   + [➡ Explore upstream dependencies interactively](#-explore-upstream-dependencies-interactively)
   + [➡ Use `go mod` directives](#-use-go-mod-directives)
   + [➡ Analyze dependencies with `goda`](#-analyze-dependencies-with-goda)
 - Code Visualization
   + [➡ Make graph of function calls in package](#-make-graph-of-function-calls-in-package)
   + [➡ Make PlantUML diagram](#-make-plantuml-diagram)
 - Assembly
   + [➡ Get assembly of Go code snippets online](#-get-assembly-of-go-code-snippets-online)
   + [➡ Get Go compiler SSA intermediary representation](#-get-go-compiler-ssa-intermediary-representation)
   + [➡ View Go assembly interactively](#-view-go-assembly-interactively)
   + [➡ Generate Go assembly in Go](#-generate-go-assembly-in-go)
 - Execute
   + [➡ Run Go function in shell](#-run-go-function-in-shell)
   + [➡ Run simple fileserver](#-run-simple-fileserver)
   + [➡ Monitor Go processes](#-monitor-go-processes)
   + [➡ Create 3D visualization of concurrency traces](#-create-3d-visualization-of-concurrency-traces)
 - Refactoring
   + [➡ Replace symbol](#-replace-symbol)
 - Errors
   + [➡ Pretty print `panic` messages](#-pretty-print-panic-messages)
 - Build
   + [➡ Manually disable or enable `cgo`](#-manually-disable-or-enable-cgo)
 - Binary
   + [➡ Make treemap breakdown of Go executable binary](#-make-treemap-breakdown-of-go-executable-binary)
 - Documentation
   + [➡ Make alternative documentation with golds](#-make-alternative-documentation-with-golds)

## Tests

### ➡ Make treemap of code coverage

Visualize distribution of code coverage in your project. This helps to identify code areas with high and low coverage. Useful when you have large project with lots of files and packages. This 2D image-hash of your project should be more representative than a single number. Also available at https://go-cover-treemap.io. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go test -coverprofile cover.out ./...
go-cover-treemap -coverprofile cover.out > out.svg
```

<div align="center"><img src="./img/hugo-code-coverage.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikolaydubina/go-cover-treemap@latest
```

### ➡ Pretty print coverage of Go code in terminal

It is similar to `go tool cover -html=cover.out` but not leaving the terminal. You can filter by functions, packages, or minimum coverage percent expressions. — [@nikandfor](https://github.com/nikandfor) / https://github.com/nikandfor/cover


```
cover
```

<div align="center"><img src="img/cover.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikandfor/cover@latest
```

### ➡ Browse code coverage of Go code in terminal

This tool lets you interactively browse Go coverage similarly to HTML version provided by official Go toolchain, but within terminal. — [@orlangure](https://github.com/orlangure) / https://github.com/orlangure/gocovsh


```
go test -cover -coverprofile coverage.out
gocovsh
gocovsh --profile profile.out
git diff --name-only | gocovsh
```

<div align="center"><img src="img/gocovsh.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/orlangure/gocovsh@latest
```

### ➡ Get packages without tests

If code coverage does not report packages without tests. This should be fast for CI. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go list -json ./... | jq -rc 'select((.TestGoFiles | length)==0) | .ImportPath'
```

Example
```
github.com/gin-gonic/gin/ginS
github.com/gin-gonic/gin/internal/json
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### ➡ Browse code coverage by file

This is very helpful tool from the official Go toolchain.


```
go test -coverprofile cover.out ./...
go tool cover -html=cover.out
```

<div align="center"><img src="./img/tool-cover-html.png" style="margin: 8px; max-height: 640px;"></div>



### ➡ Make histogram of Go files per package

Use this to see when package is too big or too small. Adjust histogram length to maximum value.


```
go list -json ./... | jq -rc '[.ImportPath, (.GoFiles | length | tostring)] | join(" ")' | perl -lane 'print (" " x (20 - $F[1]), "=" x $F[1], " ", $F[1], "\t", $F[0])'
```

Example
```
  ================== 18	github.com/gin-gonic/gin
       ============= 13	github.com/gin-gonic/gin/binding
                   = 1	github.com/gin-gonic/gin/internal/bytesconv
                   = 1	github.com/gin-gonic/gin/internal/json
         =========== 11	github.com/gin-gonic/gin/render
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### ➡ Run tests sequentially

This is in cases when you need to synchronize tests, for example in integration tests that share environment. [Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).


```
go test -p 1 -parallel 1 ./...
```


### ➡ Run tests in parallel

Add `t.Parallel` to your tests case function bodies. As per documentation, by default `-p=GOMAXPROCS` and `-parallel=GOMAXPROCS` when you run `go test`. Different packages by default run in parallel, and tests within package can be enforced to run in parallel too. Make sure to copy test case data to new variable, why explained [here](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721). [Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).

```go
    ...
    for _, tc := range tests {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            ...
```


### ➡ Run tests with pretty output

This wrapper around `go test` renders test output in easy to read format. Also supports JUnit, JSON output, skipping slow tests, running custom binary. — [@dnephin](https://github.com/dnephin) / https://github.com/gotestyourself/gotestsum


```
gotestsum --format dots
```

<div align="center"><img src="https://user-images.githubusercontent.com/442180/182284939-e08a0aa5-4504-4e30-9e88-207ef47f4537.gif" style="margin: 8px; max-height: 640px;"></div>



### ➡ Detect goroutine leaks

Refactored, tested variant of the goroutine leak detector found in both `net/http`` tests and the cockroachdb source tree. You have to call this library in your tests. — [@fortytw2](https://github.com/fortytw2) / https://github.com/fortytw2/leaktest

```go
func TestPoolContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer leaktest.CheckContext(ctx, t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}
```


## Dependencies

### ➡ Get Go version of current module

Use this in CI to setup correct Go version automatically from `go.mod`.


```
go mod edit -json | jq -r .Go
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### ➡ Get Go versions of upstream modules

Use this when upgrading version of Go or finding old modules.


```
go list -deps -json ./... | jq -rc 'select(.Standard!=true and .Module.GoVersion!=null) | [.Module.GoVersion,.Module.Path] | join(" ")' | sort -V | uniq
```

Example
```
1.11 github.com/ugorji/go/codec
1.11 golang.org/x/crypto
1.12 github.com/golang/protobuf
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### ➡ Get directly dependent modules that can be upgraded

Use this to keep your modules updated. Similar function is integrated in VSCode official Go plugin and GoLand.


```
go list -u -m $(go list -m -f '{{.Indirect}} {{.}}' all | grep '^false' | cut -d ' ' -f2) | grep '\['
```

Example
```
github.com/goccy/go-json v0.5.1 [v0.7.3]
github.com/golang/protobuf v1.3.3 [v1.5.2]
github.com/json-iterator/go v1.1.9 [v1.1.11]
```


### ➡ Get upstream modules without Go version

Use this to find outdated modules or imports that you need to upgrade.


```
go list -deps -json ./... | jq -rc 'select(.Standard!=true and .Module.GoVersion==null) | .Module.Path' | sort -u
```

Example
```
github.com/facebookgo/clock
golang.org/x/text
gopkg.in/yaml.v2
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### ➡ Get available module versions

This works even if you did not download or install module locally. This is useful to check to which version you can upgrade to, what is the latest version, and whether there are v2+ major versions recognized by Go toolchain.


```
go list -m -versions github.com/google/gofuzz
```


### ➡ Make graph of upstream modules

For each module, the node representing the greatest version (i.e., the version chosen by Go's minimal version selection algorithm) is colored green. Other nodes, which aren't in the final build list, are colored grey. — official Go team


```
go mod graph | modgraphviz | dot -Tsvg -o mod-graph.svg
```

<div align="center"><img src="./img/modgraphviz-go-featureprocessing.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
https://graphviz.org/download/
go install golang.org/x/exp/cmd/modgraphviz@latest
```

### ➡ Make graph of upstream modules with gmchart

Render in browser Go module graphs. Built with D3.js, Javascript, HTTP server in Go. — [@PaulXu-cn](https://github.com/PaulXu-cn)


```
go mod graph | gmchart
```

<div align="center"><img src="https://github.com/PaulXu-cn/go-mod-graph-chart/raw/main/show.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest
```

### ➡ Make graph of upstream packages

Use to find unexpected dependencies or visualize project. Works best for small number of packages, for large projects use `grep` to narrow down subgraph. Without `-deps` only for current module. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go list -deps -json ./... | jq -c 'select(.Standard!=true) | {from: .ImportPath, to: .Imports[]}' | jsonl-graph | dot -Tsvg > package-graph.svg
```

<div align="center"><img src="./img/packages-graph.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
https://stedolan.github.io/jq/download/
https://graphviz.org/download/
go install github.com/nikolaydubina/import-graph@latest
go install github.com/nikolaydubina/jsonl-graph@latest
```

### ➡ Scrape details about upstream modules and make graph

Use to find low quality or unmaintained dependencies. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go mod graph | import-graph -i=gomod | jsonl-graph -color-scheme=file://$PWD/basic.json | dot -Tsvg > output.svg
```

<div align="center"><img src="./img/gin-mod-graph-collected.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
https://graphviz.org/download/
go install github.com/nikolaydubina/import-graph@latest
go install github.com/nikolaydubina/jsonl-graph@latest
```

### ➡ Scrape licenses of upstream dependencies

This is tool from Google. Might be useful to collect all the licenses or checking if you can use the project for example in proprietary or commercial environment. — official Go team


```
go-licenses csv github.com/gohugoio/hugo
```

Example
```
github.com/cli/safeexec,https://github.com/cli/safeexec/blob/master/LICENSE,BSD-2-Clause
github.com/bep/tmc,https://github.com/bep/tmc/blob/master/LICENSE,MIT
github.com/aws/aws-sdk-go,https://github.com/aws/aws-sdk-go/blob/master/LICENSE.txt,Apache-2.0
github.com/jmespath/go-jmespath,https://github.com/jmespath/go-jmespath/blob/master/LICENSE,Apache-2.0
github.com/gorilla/websocket,https://github.com/gorilla/websocket/blob/master/LICENSE,BSD-2-Clause
github.com/pelletier/go-toml/v2,https://github.com/pelletier/go-toml/blob/master/v2/LICENSE,MIT
github.com/spf13/cobra,https://github.com/spf13/cobra/blob/master/LICENSE.txt,Apache-2.0
github.com/kyokomi/emoji/v2,https://github.com/kyokomi/emoji/blob/master/v2/LICENSE,MIT
go.opencensus.io,Unknown,Apache-2.0
github.com/Azure/azure-storage-blob-go/azblob,https://github.com/Azure/azure-storage-blob-go/blob/master/azblob/LICENSE,MIT
github.com/yuin/goldmark-highlighting,https://github.com/yuin/goldmark-highlighting/blob/master/LICENSE,MIT
```

Requirements
```
go install github.com/google/go-licenses@latest
```

### ➡ Explore upstream dependencies interactively

This is a tool from one of creators of Go. This tool should help explore dependencies and assist large refactorings. — [Alan Donovan](https://github.com/adonovan), official Go team / https://github.com/adonovan/spaghetti

<div align="center"><img src="https://github.com/adonovan/spaghetti/blob/main/screenshot.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/adonovan/spaghetti@latest
```

### ➡ Use `go mod` directives

Tell Go compiler which versions of upstreams to include in your build. Tell all users of your module how to deal with versions of your module.

```
// Deprecated: use example.com/mod/v2 instead.
module example.com/mod

go 1.16

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4
exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]
```


### ➡ Analyze dependencies with `goda`

This tool has extensive syntax for filtering dependencies graphs. It can work with packages and modules. — [Egon Elbre](egonelbre@gmail.com) / https://github.com/loov/goda


```
goda graph . | dot -Tsvg -o graph.svg
goda graph -cluster -short "github.com/nikolaydubina/go-cover-treemap:all" | dot -Tsvg -o graph.svg
```

<div align="center"><img src="https://github.com/loov/goda/raw/master/graph.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
https://graphviz.org/download/
go install github.com/loov/goda@latest
```

## Code Visualization

### ➡ Make graph of function calls in package

This can be helpful to quickly track which packages current package is calling and why. — [@ofabry](https://github.com/ofabry)


```
go-callvis .
```

<div align="center"><img src="https://raw.githubusercontent.com/ofabry/go-callvis/master/images/syncthing.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/ofabry/go-callvis
```

### ➡ Make PlantUML diagram

This can be useful to automatically generate visualization of classes and interfaces for go pacakges. Recommend recursive option. Render `.puml` files in for exmample [planttext.com](https://www.planttext.com). — [@bykof](https://github.com/bykof) / [github.com/bykof/go-plantuml](https://github.com/bykof/go-plantuml)


```
go-plantuml generate -d . -r -o graph.puml
```

<div align="center"><img src="https://raw.githubusercontent.com/bykof/go-plantuml/master/docs/assets/graph.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/bykof/go-plantuml@latest
```

## Assembly

### ➡ Get assembly of Go code snippets online

Use [godbolt.org](https://godbolt.org) to compile and see assembly of short Go code. You can check different platforms and compilers including `cgo`. This tool is commonly used by C++ community. — [@mattgodbolt](https://github.com/mattgodbolt)

<div align="center"><img src="./img/godbolt.png" style="margin: 8px; max-height: 640px;"></div>



### ➡ Get Go compiler SSA intermediary representation

This tool allows to check what does Go compiler do. Might be useful if you trying to optimize some code or learn more about compiler. https://golang.design/gossa. — [@changkun](https://github.com/changkun) / https://github.com/golang-design/ssaplayground

<div align="center"><img src="https://github.com/golang-design/ssaplayground/blob/main/public/assets/screen.png" style="margin: 8px; max-height: 640px;"></div>



### ➡ View Go assembly interactively

This tool lets you interactively Go statement in assembly. — [@egonelbre](https://github.com/egonelbre) / https://github.com/loov/lensm

<div align="center"><img src="https://github.com/loov/lensm/raw/main/screenshot.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install loov.dev/lensm@main
```

### ➡ Generate Go assembly in Go

Write better quality Go assembly quicker in Go language itself. This tool conveniently generates stub for Go code to call your generated assembly. Used by Go core. — [@mmcloughlin](https://github.com/mmcloughlin) / https://github.com/mmcloughlin/avo

```go
//go:build ignore
// +build ignore

package main

import . "github.com/mmcloughlin/avo/build"

func main() {
	TEXT("Add", NOSPLIT, "func(x, y uint64) uint64")
	Doc("Add adds x and y.")
	x := Load(Param("x"), GP64())
	y := Load(Param("y"), GP64())
	ADDQ(x, y)
	Store(y, ReturnIndex(0))
	RET()
	Generate()
}

```


## Execute

### ➡ Run Go function in shell

This is short and convenient for Go one-liners. This tool will print to stdout the return of a function call. — [@natefinch](https://github.com/natefinch)


```
cat README.md | gorram crypto/sha1 Sum
echo 12345 | gorram encoding/base64 StdEncoding.EncodeToString
gorram net/http Get https://google.com
```

Requirements
```
go install github.com/natefinch/gorram@latest
```

### ➡ Run simple fileserver

It takes one line to run HTTP file server in Go. Akin to famous oneliner in Python `python3 -m http.server` and `python -m SimpleHTTPServer`. Run this file as usually `go run <filename>`.

```go
package main

import "net/http"

func main() { http.ListenAndServe(":9000", http.FileServer(http.Dir("."))) }

```

<div align="center"><img src="./img/simple-fs.png" style="margin: 8px; max-height: 640px;"></div>



### ➡ Monitor Go processes

This tool from Google has lots of useful features like monitoring memory of Go processes, forcing GC, getting version of Go of process. — official Go team


```
gops
```

Example
```
983   980    uplink-soecks  go1.9   /usr/local/bin/uplink-soecks
52697 52695  gops           go1.10  /Users/jbd/bin/gops
4132  4130   foops        * go1.9   /Users/jbd/bin/foops
51130 51128  gocode         go1.9.2 /Users/jbd/bin/gocode
```

Requirements
```
go install github.com/google/gops@latest
```

### ➡ Create 3D visualization of concurrency traces

This tool creates 3D visualization of coroutines execution. There is no advanced functions and it is hard to analyze production systems. However, it could be interesting for educational purposes. — [@divan](https://github.com/divan) / https://github.com/divan/gotrace

<div align="center"><img src="https://github.com/divan/gotrace/blob/master/images/demo.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/divan/gotrace
patch Go compiler, available via Docker
more instructions in original repo
```

## Refactoring

### ➡ Replace symbol

I found this in announcement [notice](https://github.com/golang/go/commit/2580d0e08d5e9f979b943758d3c49877fb2324cb) of Go 1.18 for changes to `interface{}` to `any`. This can be useful for other refactorings too.


```
gofmt -w -r 'interface{} -> any' .
```


## Errors

### ➡ Pretty print `panic` messages

This tool will be useful for reading `panic` messages. Need to redirect STDERR to this tool with `panic` stack traces. The tool has HTML output and does lots of deduplication and enhancements. Refer to examples in original repo.


```
go test -v |& pp
```

<div align="center"><img src="https://raw.githubusercontent.com/wiki/maruel/panicparse/parse.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/maruel/panicparse/v2/cmd/pp@latest
```

## Build

### ➡ Manually disable or enable `cgo`

Disable `cgo` with `CGO_ENABLED=0` and enable with `CGO_ENABLED=1`. If you don't, `cgo` may end-up being enabled or code dynamically linked if, for example, you use some `net` or `os` packages. You may want to disable `cgo` to improve performance, since complier and runtime would have easier job optimizing code. This also should reduce your image size, as you can have alpine image with less shared libraries.


## Binary

### ➡ Make treemap breakdown of Go executable binary

This can be useful for studying Go compiler, large projects, projects with C/C++ and `cgo`, 3rd party dependencies, embedding. However, total size may not be something to worry about for your executable. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
```

<div align="center"><img src="https://github.com/nikolaydubina/go-binsize-treemap/blob/main/docs/hugo.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikolaydubina/go-binsize-treemap@latest
```

## Documentation

### ➡ Make alternative documentation with golds

It has additional information like implementations of interface; promoted methods. The tool has nice minimalistic aesthetics. — [Tapir Liu](https://www.tapirgames.com) / https://github.com/go101/golds


```
golds ./...
```

<div align="center"><img src="img/golds.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install go101.org/golds@latest
```


