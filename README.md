<h1 align="center">Go Recipes 🦩 </h1>
<p align="center">Handy well-known and <i>lesser</i>-known tools for Go projects</p>

> _Know some cool tool or one-liner? Have a feature request or an idea?_  
> _Feel free to edit this page or create an Issue!_  

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fnikolaydubina%2Fgo-recipes&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)
[![go-recipes](https://raw.githubusercontent.com/nikolaydubina/go-recipes/main/badge.svg?raw=true)](https://github.com/nikolaydubina/go-recipes)

## Contents

 - Testing
   + [➡ Make treemap of code coverage with `go-cover-treemap`](#-make-treemap-of-code-coverage-with-go-cover-treemap)
   + [➡ Browse coverage in browser with `gocov-html`](#-browse-coverage-in-browser-with-gocov-html)
   + [➡ Browse code coverage by file](#-browse-code-coverage-by-file)
   + [➡ Browse code coverage in terminal with `gocovsh`](#-browse-code-coverage-in-terminal-with-gocovsh)
   + [➡ Pretty print coverage in terminal with `nikandfor/cover`](#-pretty-print-coverage-in-terminal-with-nikandforcover)
   + [➡ Run tests sequentially](#-run-tests-sequentially)
   + [➡ Run tests in parallel](#-run-tests-in-parallel)
   + [➡ Detect goroutine leaks with `leaktest`](#-detect-goroutine-leaks-with-leaktest)
   + [➡ Run tests with pretty output with `gotestsum`](#-run-tests-with-pretty-output-with-gotestsum)
   + [➡ Summarize `go test` with `tparse`](#-summarize-go-test-with-tparse)
   + [➡ Colorize and decorate `go test` with `richgo`](#-colorize-and-decorate-go-test-with-richgo)
   + [➡ Colorize `go test` with `gotest`](#-colorize-go-test-with-gotest)
   + [➡ Get packages without tests](#-get-packages-without-tests)
   + [➡ Perform Mutation Testing with `ooze`](#-perform-mutation-testing-with-ooze)
   + [➡ Perform Mutation Testing with `avito-tech/go-mutesting`](#-perform-mutation-testing-with-avito-techgo-mutesting)
   + [➡ Perform Mutation Testing with `go-mutesting`](#-perform-mutation-testing-with-go-mutesting)
   + [➡ Trace tests with `go-test-trace`](#-trace-tests-with-go-test-trace)
 - Dependencies
   + [➡ Get Go version of current module](#-get-go-version-of-current-module)
   + [➡ Get Go versions of upstream modules](#-get-go-versions-of-upstream-modules)
   + [➡ Get directly dependent modules that can be upgraded](#-get-directly-dependent-modules-that-can-be-upgraded)
   + [➡ Get upstream modules without Go version](#-get-upstream-modules-without-go-version)
   + [➡ Get available module versions](#-get-available-module-versions)
   + [➡ Make graph of upstream modules with `modgraphviz`](#-make-graph-of-upstream-modules-with-modgraphviz)
   + [➡ Make graph of upstream modules with `gmchart`](#-make-graph-of-upstream-modules-with-gmchart)
   + [➡ Make graph of upstream packages with `import-graph`](#-make-graph-of-upstream-packages-with-import-graph)
   + [➡ Scrape details about upstream modules and make graph with `import-graph`](#-scrape-details-about-upstream-modules-and-make-graph-with-import-graph)
   + [➡ Scrape licenses of upstream dependencies with `go-licenses`](#-scrape-licenses-of-upstream-dependencies-with-go-licenses)
   + [➡ Explore upstream dependencies interactively with `spaghetti`](#-explore-upstream-dependencies-interactively-with-spaghetti)
   + [➡ Use `go mod` directives](#-use-go-mod-directives)
   + [➡ Analyze dependencies with `goda`](#-analyze-dependencies-with-goda)
 - Code Visualization
   + [➡ Make C4 diagram with `go-structurizr`](#-make-c4-diagram-with-go-structurizr)
   + [➡ Make graph of function calls with `callgraph`](#-make-graph-of-function-calls-with-callgraph)
   + [➡ Make graph of function calls in package with `go-callvis`](#-make-graph-of-function-calls-in-package-with-go-callvis)
   + [➡ Make PlantUML diagram with `goplantuml`](#-make-plantuml-diagram-with-goplantuml)
   + [➡ Make PlantUML diagram with `go-plantuml`](#-make-plantuml-diagram-with-go-plantuml)
   + [➡ Make 3D chart of Go codebase with `gocity`](#-make-3d-chart-of-go-codebase-with-gocity)
   + [➡ Make histogram of Go files per package](#-make-histogram-of-go-files-per-package)
   + [➡ Explore Go code in browser powered by `go-guru` with `pythia`](#-explore-go-code-in-browser-powered-by-go-guru-with-pythia)
   + [➡ (archived) Interactively visualize packages with `goexplorer`](#-archived-interactively-visualize-packages-with-goexplorer)
 - Static Analysis
   + [➡ Run default static analysis with `go vet`](#-run-default-static-analysis-with-go-vet)
   + [➡ Run custom static analysis tool with `go vet`](#-run-custom-static-analysis-tool-with-go-vet)
   + [➡ Run official static analyzers not included in `go vet`](#-run-official-static-analyzers-not-included-in-go-vet)
   + [➡ Detect most common issues with `staticcheck`](#-detect-most-common-issues-with-staticcheck)
   + [➡ Detect non-exhaustive switch and map with `exhaustive`](#-detect-non-exhaustive-switch-and-map-with-exhaustive)
   + [➡ Detect structs with uninitialized fields with `go-exhaustruct`](#-detect-structs-with-uninitialized-fields-with-go-exhaustruct)
   + [➡ Detect unsafe code with `go-safer`](#-detect-unsafe-code-with-go-safer)
   + [➡ Detect unnecessary type conversions with `unconvert`](#-detect-unnecessary-type-conversions-with-unconvert)
   + [➡ Detect global variables with `gochecknoglobals`](#-detect-global-variables-with-gochecknoglobals)
   + [➡ Detect slices that could be preallocated with `prealloc`](#-detect-slices-that-could-be-preallocated-with-prealloc)
   + [➡ Detect unnecessary import aliases with `unimport`](#-detect-unnecessary-import-aliases-with-unimport)
   + [➡ Detect naked returns with `nakedret`](#-detect-naked-returns-with-nakedret)
   + [➡ Detect mixing pointer and value method receivers with `smrcptr`](#-detect-mixing-pointer-and-value-method-receivers-with-smrcptr)
   + [➡ Detect vertical function ordering with `vertfn`](#-detect-vertical-function-ordering-with-vertfn)
   + [➡ Calculate Cognitive Complexity with `gocognit`](#-calculate-cognitive-complexity-with-gocognit)
   + [➡ Calculate Cyclomatic Complexity with `gocyclo`](#-calculate-cyclomatic-complexity-with-gocyclo)
   + [➡ Calculate age of comments with `go-commentage`](#-calculate-age-of-comments-with-go-commentage)
   + [➡ (archived) Ensure `if` statements using short assignment with `ifshort`](#-archived-ensure-if-statements-using-short-assignment-with-ifshort)
   + [➡ Perform Taint Analysis with `taint`](#-perform-taint-analysis-with-taint)
   + [➡ Visualize struct layout with `structlayout`](#-visualize-struct-layout-with-structlayout)
   + [➡ Rely on compiler for stricter Enums](#-rely-on-compiler-for-stricter-enums)
 - Code Generation
   + [➡ Run `go:generate` in parallel](#-run-gogenerate-in-parallel)
   + [➡ Generate `String` method for enum types](#-generate-string-method-for-enum-types)
   + [➡ Modify struct field tags with `gomodifytags`](#-modify-struct-field-tags-with-gomodifytags)
   + [➡ Generate Table Driven Tests with `gotests`](#-generate-table-driven-tests-with-gotests)
 - Refactoring
   + [➡ Replace symbol with `gofmt`](#-replace-symbol-with-gofmt)
   + [➡ Keep consistent ordering of imports with `gci`](#-keep-consistent-ordering-of-imports-with-gci)
   + [➡ Keep consistent ordering of imports with `goimportx`](#-keep-consistent-ordering-of-imports-with-goimportx)
 - Errors
   + [➡ Pretty print `panic` messages with `panicparse`](#-pretty-print-panic-messages-with-panicparse)
 - Build
   + [➡ Show compiler optimization decisions on heap and inlining](#-show-compiler-optimization-decisions-on-heap-and-inlining)
   + [➡ Disable inlining](#-disable-inlining)
   + [➡ Aggressive inlining](#-aggressive-inlining)
   + [➡ Manually disable or enable `cgo`](#-manually-disable-or-enable-cgo)
   + [➡ Include metadata in binary during compilation with `ldflags`](#-include-metadata-in-binary-during-compilation-with-ldflags)
   + [➡ Make treemap breakdown of Go executable binary with `go-binsize-treemap`](#-make-treemap-breakdown-of-go-executable-binary-with-go-binsize-treemap)
   + [➡ Custom import path](#-custom-import-path)
   + [➡ Custom import path with `govanityurls`](#-custom-import-path-with-govanityurls)
   + [➡ Custom import path with `sally`](#-custom-import-path-with-sally)
   + [➡ Custom import path with `kkn.fi/vanity`](#-custom-import-path-with-kknfivanity)
   + [➡ Custom import path enforcement](#-custom-import-path-enforcement)
 - Assembly
   + [➡ Get assembly of Go code snippets online](#-get-assembly-of-go-code-snippets-online)
   + [➡ Get Go SSA intermediary representation with `ssaplayground`](#-get-go-ssa-intermediary-representation-with-ssaplayground)
   + [➡ View Go assembly interactively with `lensm`](#-view-go-assembly-interactively-with-lensm)
   + [➡ Generate Go assembly in Go with `avo`](#-generate-go-assembly-in-go-with-avo)
   + [➡ Generate AST for code snippets](#-generate-ast-for-code-snippets)
   + [➡ Visualize Go SSA function using Graphviz with `go-ssaviz`](#-visualize-go-ssa-function-using-graphviz-with-go-ssaviz)
   + [➡ (archived) Make graph of AST with `astgraph`](#-archived-make-graph-of-ast-with-astgraph)
 - Execution
   + [➡ Run interactive Go interpreter with `gomacro`](#-run-interactive-go-interpreter-with-gomacro)
   + [➡ Run Go function in shell with `gorram`](#-run-go-function-in-shell-with-gorram)
   + [➡ Run simple fileserver](#-run-simple-fileserver)
   + [➡ Create 3D visualization of concurrency traces with `gotrace`](#-create-3d-visualization-of-concurrency-traces-with-gotrace)
 - Monitoring
   + [➡ Monitor goroutines with `grmon`](#-monitor-goroutines-with-grmon)
   + [➡ Monitor Go processes with `gops`](#-monitor-go-processes-with-gops)
   + [➡ Visualise Go runtime metrics in browser with `statsviz`](#-visualise-go-runtime-metrics-in-browser-with-statsviz)
   + [➡ Auto-Instrument all functions with `go-instrument`](#-auto-instrument-all-functions-with-go-instrument)
   + [➡ Auto-Instrument all functions with `otelinji`](#-auto-instrument-all-functions-with-otelinji)
 - Benchmarking
   + [➡ Run benchmarks](#-run-benchmarks)
   + [➡ Table-driven benchmarks](#-table-driven-benchmarks)
   + [➡ Generate benchmak CPU and Memory profiles](#-generate-benchmak-cpu-and-memory-profiles)
   + [➡ Visualize callgraph of profiles with `pprof`](#-visualize-callgraph-of-profiles-with-pprof)
   + [➡ Visualize flamegraphs of profiles with `pprof`](#-visualize-flamegraphs-of-profiles-with-pprof)
   + [➡ Visualize profiles online](#-visualize-profiles-online)
   + [➡ Get delta between two benchmarks with `benchstat`](#-get-delta-between-two-benchmarks-with-benchstat)
   + [➡ Get summary of benchmarks with `benchstat`](#-get-summary-of-benchmarks-with-benchstat)
   + [➡ Continuous benchmarking](#-continuous-benchmarking)
   + [➡ Continuous benchmarking with `gobenchdata`](#-continuous-benchmarking-with-gobenchdata)
   + [➡ Continuous benchmarking with `benchdiff`](#-continuous-benchmarking-with-benchdiff)
   + [➡ Continuous benchmarking with `cob`](#-continuous-benchmarking-with-cob)
   + [➡ Generate live traces with `net/http/trace`](#-generate-live-traces-with-nethttptrace)
   + [➡ Generate traces with `go test`](#-generate-traces-with-go-test)
   + [➡ View traces with `go tool trace`](#-view-traces-with-go-tool-trace)
   + [➡ Get wallclock traces with `fgtrace`](#-get-wallclock-traces-with-fgtrace)
   + [➡ Get on/off CPU profiles with `fgprof`](#-get-onoff-cpu-profiles-with-fgprof)
 - Documentation
   + [➡ Make alternative documentation with `golds`](#-make-alternative-documentation-with-golds)
   + [➡ Read Go binary documentation in `man` format with `goman`](#-read-go-binary-documentation-in-man-format-with-goman)
 - Style Guide
   + [➡ Google](#style-guide)
   + [➡ Uber](#style-guide)
   + [➡ Go Code Review Comments](#style-guide)

## Testing

### [⏫](#contents)➡ Make treemap of code coverage with [go-cover-treemap](https://github.com/nikolaydubina/go-cover-treemap)

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

### [⏫](#contents)➡ Browse coverage in browser with [gocov-html](https://github.com/matm/gocov-html)

Browse coverage in statically generated HTML page. Multiple styles are supported. You may need to convert coverage report into `gocov` format. — [@matm](https://github.com/matm)


```
gocov test strings | gocov-html -t golang > strings.html
gocov test encoding/csv strings | gocov-html -t kit > strings.html
```

<div align="center"><img src="./img/gocov-html.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/axw/gocov/gocov@latest
go install github.com/matm/gocov-html/cmd/gocov-html@latest
```

### [⏫](#contents)➡ Browse code coverage by file

This is very helpful tool from the official Go toolchain. Similar visualization is integrated into VSCode and Goland, but can be used separately.


```
go test -coverprofile cover.out ./...
go tool cover -html=cover.out
```

<div align="center"><img src="./img/tool-cover-html.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Browse code coverage in terminal with [gocovsh](https://github.com/orlangure/gocovsh)

Interactively browse Go code coverage similarly to HTML provided by official Go toolchain, but in terminal. — [@orlangure](https://github.com/orlangure)


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

### [⏫](#contents)➡ Pretty print coverage in terminal with [nikandfor/cover](https://github.com/nikandfor/cover)

It is similar to `go tool cover -html=cover.out` but in terminal. You can filter by functions, packages, minimum coverage, and more. — [@nikandfor](https://github.com/nikandfor)


```
cover
```

<div align="center"><img src="img/cover.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikandfor/cover@latest
```

### [⏫](#contents)➡ Run tests sequentially

Use when you need to synchronize tests, for example in integration tests that share environment. [Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).


```
go test -p 1 -parallel 1 ./...
```


### [⏫](#contents)➡ Run tests in parallel

Add `t.Parallel` to your tests case function bodies. As per documentation, by default `-p=GOMAXPROCS` and `-parallel=GOMAXPROCS` when you run `go test`. Different packages by default run in parallel, and tests within package can be enforced to run in parallel too. Make sure to copy test case data to new variable, why explained [here](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721). [Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).

```go
...
for _, tc := range tests {
    tc := tc
    t.Run(tc.name, func(t *testing.T) {
        t.Parallel()
        ...
```


### [⏫](#contents)➡ Detect goroutine leaks with [leaktest](https://github.com/fortytw2/leaktest)

Refactored, tested variant of the goroutine leak detector found in both `net/http` tests and the cockroachdb source tree. You have to call this library in your tests. — [@fortytw2](https://github.com/fortytw2)

```
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


### [⏫](#contents)➡ Run tests with pretty output with [gotestsum](https://github.com/gotestyourself/gotestsum)

This wrapper around `go test` renders test output in easy to read format. Also supports JUnit, JSON output, skipping slow tests, running custom binary. — [@dnephin](https://github.com/dnephin)


```
gotestsum --format dots
```

<div align="center"><img src="https://user-images.githubusercontent.com/442180/182284939-e08a0aa5-4504-4e30-9e88-207ef47f4537.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install gotest.tools/gotestsum@latest
```

### [⏫](#contents)➡ Summarize `go test` with [tparse](https://github.com/mfridman/tparse)

This lightweight wrapper around STDOUT of JSON of `go test` will nicely render colorized test status, details of failures, duration, coverage, and package summary. — [@mfridman](https://github.com/mfridman)


```
set -o pipefail && go test ./... -json | tparse -all
```

<div align="center"><img src="./img/tparse.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/mfridman/tparse@latest
```

### [⏫](#contents)➡ Colorize and decorate `go test` with [richgo](https://github.com/kyoh86/richgo)

Add colors and enrich `go test` output. It can be used in CI pipeline and has lots of alternative visualizations and options. — [@kyoh86](https://github.com/kyoh86)


```
richgo test ./...
```

<div align="center"><img src="https://asciinema.org/a/99810.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/kyoh86/richgo@latest
```

### [⏫](#contents)➡ Colorize `go test` with [gotest](https://github.com/rakyll/gotest)

Add colors to `go test` output. Very lightweight wrapper around `go test` STDOUT. — [@rakyll](https://github.com/rakyll)


```
gotest ./...
```

<div align="center"><img src="https://raw.githubusercontent.com/jonasbn/go-test-demo/1.0.0/gotest-go-test-demo.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/rakyll/gotest@latest
```

### [⏫](#contents)➡ Get packages without tests

If code coverage does not report packages without tests. For example for CI or quality control.


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

### [⏫](#contents)➡ Perform Mutation Testing with [ooze](https://github.com/gtramontina/ooze)

Mutation testing is a technique used to assess the quality and coverage of test suites. It involves introducing controlled changes to the code base, simulating common programming mistakes. These changes are, then, put to test against the test suites. A failing test suite is a good sign. It indicates that the tests are identifying mutations in the code—it "killed the mutant". If all tests pass, we have a surviving mutant. This highlights an area with weak coverage. It is an opportunity for improvement. — [@gtramontina](https://github.com/gtramontina)


```
go test -v -tags=mutation
```

<div align="center"><img src="https://github.com/gtramontina/ooze/blob/main/.assets/report.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go get github.com/gtramontina/ooze
```

### [⏫](#contents)➡ Perform Mutation Testing with [avito-tech/go-mutesting](https://github.com/avito-tech/go-mutesting)

This is fork of [zimmski/go-mutesting](https://github.com/zimmski/go-mutesting). It has more mutators and latest updates. — [@vasiliyyudin](https://github.com/vasiliyyudin)


```
go-mutesting ./...
```

```go
for _, d := range opts.Mutator.DisableMutators {
  pattern := strings.HasSuffix(d, "*")

-	if (pattern && strings.HasPrefix(name, d[:len(d)-2])) || (!pattern && name == d) {
+	if (pattern && strings.HasPrefix(name, d[:len(d)-2])) || false {
    continue MUTATOR
  }
}
```

Requirements
```
go install github.com/avito-tech/go-mutesting/cmd/go-mutesting@latest
```

### [⏫](#contents)➡ Perform Mutation Testing with [go-mutesting](https://github.com/zimmski/go-mutesting)

Find common bugs source code that would pass tests. This is earliest tool for mutation testing in Go. More functions and permutations were added in other mutation Go tools it inspired. — [@zimmski](https://github.com/zimmski)


```
go-mutesting ./...
```

```go
for _, d := range opts.Mutator.DisableMutators {
  pattern := strings.HasSuffix(d, "*")

-	if (pattern && strings.HasPrefix(name, d[:len(d)-2])) || (!pattern && name == d) {
+	if (pattern && strings.HasPrefix(name, d[:len(d)-2])) || false {
    continue MUTATOR
  }
}
```

Requirements
```
go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest
```

### [⏫](#contents)➡ Trace tests with [go-test-trace](https://github.com/rakyll/go-test-trace)

Collect test execution as distributed traces. This is useful for tracking test duration, failures, flakiness. You distributed tracing storage, search, UI, exploration, dashboards, alarms — all will automatically become test status collection. If you run integration tests in your CI, then it is particularly handy to investigate your integration tests same way as real requests, such as Go processes, databases, etc. However, if you do not have distributed traces, it is still useful for adhoc investigations. This tool processes STDOUT of `go test`. No automatic instrumentation is done. — [@rakyll](https://github.com/rakyll)


```
go-test-trace ./...
```

<div align="center"><img src="https://camo.githubusercontent.com/1bbb99d14634e097828aff76e17427c0d834b2b37b7ef6c4b15ad01e5b7ac526/68747470733a2f2f692e696d6775722e636f6d2f45313850596b342e706e67" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
# open telemetry collector
# traces UI (Datadog, Jaeger, Honeycomb, NewRelic)
go install github.com/rakyll/go-test-trace@latest
```

## Dependencies

### [⏫](#contents)➡ Get Go version of current module

For example, setup correct Go version automatically from `go.mod` in CI.


```
go mod edit -json | jq -r .Go
```

Requirements
```
https://stedolan.github.io/jq/download/
```

### [⏫](#contents)➡ Get Go versions of upstream modules

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

### [⏫](#contents)➡ Get directly dependent modules that can be upgraded

Keep your modules updated. Similar function is integrated in VSCode official Go plugin and GoLand.


```
go list -u -m $(go list -m -f '{{.Indirect}} {{.}}' all | grep '^false' | cut -d ' ' -f2) | grep '\['
```

Example
```
github.com/goccy/go-json v0.5.1 [v0.7.3]
github.com/golang/protobuf v1.3.3 [v1.5.2]
github.com/json-iterator/go v1.1.9 [v1.1.11]
```


### [⏫](#contents)➡ Get upstream modules without Go version

Find outdated modules or imports that you need to upgrade.


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

### [⏫](#contents)➡ Get available module versions

This works even if you did not download or install module locally. This is useful to check to which version you can upgrade to, what is the latest version, and whether there are v2+ major versions recognized by Go toolchain.


```
go list -m -versions github.com/google/gofuzz
```


### [⏫](#contents)➡ Make graph of upstream modules with [modgraphviz](https://golang.org/x/exp/cmd/modgraphviz)

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

### [⏫](#contents)➡ Make graph of upstream modules with [gmchart](https://github.com/PaulXu-cn/go-mod-graph-chart/gmchart)

Render in browser Go module graphs. Built with D3.js, Javascript, HTTP server in Go. — [@PaulXu-cn](https://github.com/PaulXu-cn)


```
go mod graph | gmchart
```

<div align="center"><img src="https://github.com/PaulXu-cn/go-mod-graph-chart/raw/main/show.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest
```

### [⏫](#contents)➡ Make graph of upstream packages with [import-graph](https://github.com/nikolaydubina/import-graph)

Find unexpected dependencies or visualize project. Works best for small number of packages, for large projects use `grep` to narrow down subgraph. Without `-deps` only for current module. — [@nikolaydubina](https://github.com/nikolaydubina)


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

### [⏫](#contents)➡ Scrape details about upstream modules and make graph with [import-graph](https://github.com/nikolaydubina/import-graph)

Find low quality or unmaintained dependencies. — [@nikolaydubina](https://github.com/nikolaydubina)


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

### [⏫](#contents)➡ Scrape licenses of upstream dependencies with [go-licenses](https://github.com/google/go-licenses)

Collect all the licenses for checking if you can use the project, for example in proprietary or commercial environment. — Google


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

### [⏫](#contents)➡ Explore upstream dependencies interactively with [spaghetti](https://github.com/adonovan/spaghetti)

Useful in large refactorings, dependency breaking, physical layout changes. — [Alan Donovan](https://github.com/adonovan), official Go team

<div align="center"><img src="https://github.com/adonovan/spaghetti/blob/main/screenshot.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/adonovan/spaghetti@latest
```

### [⏫](#contents)➡ Use `go mod` directives

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


### [⏫](#contents)➡ Analyze dependencies with [goda](https://github.com/loov/goda)

This tool has extensive syntax for filtering dependencies graphs. It can work with packages and modules. — [Egon Elbre](egonelbre@gmail.com)


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

### [⏫](#contents)➡ Make C4 diagram with [go-structurizr](https://github.com/krzysztofreczek/go-structurizr)

This library provides tools to generate [C4](https://c4model.com) diagrams. The process is a bit involved, however you get diagram generated from real Go code automatically. Steps are outlined in [blog](https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/). — [@krzysztofreczek](https://github.com/krzysztofreczek)

<div align="center"><img src="https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/tdl-go_structurizr_1_2.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
manually defining Go main.go script to invoke library
graphviz
manual coloring spec (DB, calsses)
```

### [⏫](#contents)➡ Make graph of function calls with [callgraph](https://golang.org/x/tools/cmd/callgraph)

Visualize complex or new project quickly or to study project. Requires `main.go` in module. Supports Graphviz output format. Has many options for filtering and formatting. — official Go team


```
callgraph -format graphviz . | dot -Tsvg -o graph.svg
recommend: grep <package/class/func of interest>
recommend: grep -v Error since many packages report error
recommend: adding `rankdir=LR;` to graphviz file for denser graph
recommend: you would have to manually fix graphviz file first and last line
```

<div align="center"><img src="img/callgraph.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install golang.org/x/tools/cmd/callgraph@latest
```

### [⏫](#contents)➡ Make graph of function calls in package with [go-callvis](https://github.com/ofabry/go-callvis)

Quickly track which packages current package is calling and why. — [@ofabry](https://github.com/ofabry)


```
go-callvis .
```

<div align="center"><img src="https://raw.githubusercontent.com/ofabry/go-callvis/master/images/syncthing.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/ofabry/go-callvis
```

### [⏫](#contents)➡ Make PlantUML diagram with [goplantuml](https://github.com/jfeliu007/goplantuml)

Generates class diagram in widely used format with the information on structs, interfaces and their relationships. Render `.puml` files in for example [planttext.com](https://www.planttext.com). — [@jfeliu007](https://github.com/jfeliu007)


```
goplantuml -recursive path/to/gofiles path/to/gofiles2
```

<div align="center"><img src="https://raw.githubusercontent.com/jfeliu007/goplantuml/master/example/example.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go get github.com/jfeliu007/goplantuml/parser
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
```

### [⏫](#contents)➡ Make PlantUML diagram with [go-plantuml](https://github.com/bykof/go-plantuml)

Automatically generate visualization of classes and interfaces for go packages. Recommend recursive option. Render `.puml` files in for example [planttext.com](https://www.planttext.com). — [@bykof](https://github.com/bykof)


```
go-plantuml generate -d . -r -o graph.puml
```

<div align="center"><img src="https://raw.githubusercontent.com/bykof/go-plantuml/master/docs/assets/graph.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/bykof/go-plantuml@latest
```

### [⏫](#contents)➡ Make 3D chart of Go codebase with [gocity](https://github.com/rodrigo-brito/gocity)

Fresh artistic perspective on Go codebase. `GoCity` is an implementation of the Code City metaphor for visualizing source code - folders are districts; files are buildings; structs are buildings on the top of their files. This project has research paper "[GoCity Code City for Go](https://homepages.dcc.ufmg.br/~mtov/pub/2019-saner-gocity.pdf)" at SANER'19. Also available at [go-city.github.io](https://go-city.github.io). — [@rodrigo-brito](https://github.com/rodrigo-brito)

<div align="center"><img src="img/gocity.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/rodrigo-brito/gocity@latest
```

### [⏫](#contents)➡ Make histogram of Go files per package

Find when package is too big or too small. Adjust histogram length to maximum value.


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

### [⏫](#contents)➡ Explore Go code in browser powered by `go-guru` with [pythia](https://github.com/fzipp/pythia)

Explore Go source code in browser. It provides exported symbols summary for navigation. It answers questions like: definition; callers; implementers. It is browser frontend based on [go-guru](https://docs.google.com/document/d/1_Y9xCEMj5S-7rv2ooHpZNH15JgRT5iM742gJkw5LtmQ/edit), which was developed by Go core team from Google. — [@fzipp](https://github.com/fzipp)


```
pythia net/http
```

<div align="center"><img src="https://camo.githubusercontent.com/a7baec2bada145869272edf97d1123d1717ed68922c159b027d261bd6e1faeff/68747470733a2f2f7261772e6769746875622e636f6d2f667a6970702f7079746869612f67682d70616765732f696d616765732f7079746869615f73637265656e73686f742e706e67" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/fzipp/pythia@latest
go install golang.org/x/tools/cmd/guru@latest
```

### [⏫](#contents)➡ (archived) Interactively visualize packages with [goexplorer](https://github.com/ofabry/goexplorer)

Based on `go-callvis`, this tool is an interactive package explorer of packages. This tool have not been updated for a long time. — [@ofabry](https://github.com/ofabry)

<div align="center"><img src="https://github.com/ofabry/goexplorer/raw/master/images/screen.png" style="margin: 8px; max-height: 640px;"></div>



## Static Analysis

### [⏫](#contents)➡ Run default static analysis with `go vet`

Official tool for static analysis of Go programs, with 27+ static analyzers. — official Go team


```
go vet ./...
```


### [⏫](#contents)➡ Run custom static analysis tool with `go vet`

Standard `go vet` can be used to run custom analyzers binaries. Third party analyzers are supported. Lots of official analyzers not included by default into `go vet`. Analyzer has to satisfy interface and command described here https://pkg.go.dev/golang.org/x/tools/go/analysis. Refer for https://pkg.go.dev/golang.org/x/tools/go/analysis/passes for full list of official Go analyzers. — official Go team


```
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
go vet -vettool=$(which shadow)
```


### [⏫](#contents)➡ Run official static analyzers not included in `go vet`

There are many analyzers not included in `go vet`. These tools are experimental and may not work as expected (e.g. `usesgenerics` does not work). Refer to for full list https://pkg.go.dev/golang.org/x/tools/go/analysis. — official Go team

```go
package main

import (
  "golang.org/x/tools/go/analysis/multichecker"

  "golang.org/x/tools/go/analysis/passes/atomicalign"
  "golang.org/x/tools/go/analysis/passes/deepequalerrors"
  "golang.org/x/tools/go/analysis/passes/fieldalignment"
  "golang.org/x/tools/go/analysis/passes/nilness"
  "golang.org/x/tools/go/analysis/passes/reflectvaluecompare"
  "golang.org/x/tools/go/analysis/passes/shadow"
  "golang.org/x/tools/go/analysis/passes/sortslice"
  "golang.org/x/tools/go/analysis/passes/unusedwrite"
  "golang.org/x/tools/go/analysis/passes/usesgenerics"
)

func main() {
  multichecker.Main(
    atomicalign.Analyzer,         // checks for non-64-bit-aligned arguments to sync/atomic functions
    deepequalerrors.Analyzer,     // checks for the use of reflect.DeepEqual with error values
    fieldalignment.Analyzer,      // detects structs that would use less memory if their fields were sorted
    nilness.Analyzer,             // inspects the control-flow graph of an SSA function and reports errors such as nil pointer dereferences and degenerate nil pointer comparisons
    reflectvaluecompare.Analyzer, // checks for accidentally using == or reflect.DeepEqual to compare reflect.Value values
    shadow.Analyzer,              // checks for shadowed variables
    sortslice.Analyzer,           // checks for calls to sort.Slice that do not use a slice type as first argument
    unusedwrite.Analyzer,         // checks for unused writes to the elements of a struct or array object
    usesgenerics.Analyzer,        // checks for usage of generic features added in Go 1.18
  )
}
```


### [⏫](#contents)➡ Detect most common issues with [staticcheck](https://github.com/dominikh/go-tools)

Start custom linters with this well-known linter. It contains 150+ high quality low false positive rate linters. It is widely adopted by Open Source and tech companies. [staticcheck.io](https://staticcheck.io/). — [@dominikh](https://github.com/dominikh)


```
staticcheck ./...
```

Requirements
```
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### [⏫](#contents)➡ Detect non-exhaustive switch and map with [exhaustive](https://github.com/nishanths/exhaustive)

This `go vet` compatible analyzer checks for exhaustive switch statemnts and map literals. It works for enums with underyling integer, float, or string types (struct based enums are not supported). — [@nishanths](https://github.com/nishanths)


```
exhaustive ./...
```

```go
package token

type Token int

const (
  Add Token = iota
  Subtract
  Multiply
  Quotient
  Remainder
)

package calc

import "token"

func f(t token.Token) {
  switch t {
  case token.Add:
  case token.Subtract:
  case token.Multiply:
  default:
  }
}

func g(t token.Token) string {
  return map[token.Token]string{
    token.Add:      "add",
    token.Subtract: "subtract",
    token.Multiply: "multiply",
  }[t]
}
```

Example
```
calc.go:6:2: missing cases in switch of type token.Token: Quotient, Remainder
calc.go:15:9: missing map keys of type token.Token: Quotient, Remainder
```

Requirements
```
go install github.com/nishanths/exhaustive/cmd/exhaustive@latest
```

### [⏫](#contents)➡ Detect structs with uninitialized fields with [go-exhaustruct](https://github.com/GaijinEntertainment/go-exhaustruct)

This tool finds instatiations of structs with zero values. It supports struct tags to mark fields as optional. This may help to prevent unexpected zero values. — [@xobotyi](https://github.com/xobotyi)


```
exhaustruct ./...
```

```go
type Shape struct {
  Length int
  Width  int
  volume    int
  Perimeter int `exhaustruct:"optional"`
}

// valid
var a Shape = Shape{
  Length: 5,
  Width:  3,
  volume: 5,
}

// invalid, `volume` is missing
var b Shape = Shape{
  Length: 5,
  Width:  3,
}
```

Requirements
```
go get -u github.com/GaijinEntertainment/go-exhaustruct/v3/cmd/exhaustruct
```

### [⏫](#contents)➡ Detect unsafe code with [go-safer](https://github.com/jlauinger/go-safer)

Find incorrect uses of `reflect.SliceHeader`, `reflect.StringHeader`, and unsafe casts between structs with architecture-sized fields. Reseach paper ["Uncovering the Hidden Dangers Finding Unsafe Go Code in the Wild"](https://arxiv.org/abs/2010.11242) presented at 19th IEEE International Conference on Trust, Security and Privacy in Computing and Communications (TrustCom 2020). — [@jlauinger](https://github.com/jlauinger)


```
go-safer ./...
```

Example
```
# github.com/jlauinger/go-safer/passes/sliceheader/testdata/src/bad/composite_literal
composite_literal/composite_literal.go:10:9: reflect header composite literal found
composite_literal/composite_literal.go:10:9: reflect header composite literal found
# github.com/jlauinger/go-safer/passes/sliceheader/testdata/src/bad/header_in_struct
header_in_struct/header_in_struct.go:16:2: assigning to reflect header object
```

Requirements
```
go install github.com/jlauinger/go-safer@latest
```

### [⏫](#contents)➡ Detect unnecessary type conversions with [unconvert](https://github.com/mdempsky/unconvert)

Identify expressions like `T(x)` where `x` is already has type `T`. This tool can identify conversions that force intermediate rounding. It also can overwrite files with fix. This tool is not using `golang.org/x/tools/go/analysis` toolchain. — [@mdempsky](https://github.com/mdempsky)


```
unconvert ./...
```

```
$ unconvert -v bytes fmt
GOROOT/src/bytes/reader.go:117:14: unnecessary conversion
                abs = int64(r.i) + offset
                          ^
GOROOT/src/fmt/print.go:411:21: unnecessary conversion
        p.fmt.integer(int64(v), 16, unsigned, udigits)
                          ^
```

Requirements
```
go install github.com/mdempsky/unconvert@latest
```

### [⏫](#contents)➡ Detect global variables with [gochecknoglobals](https://github.com/leighmcculloch/gochecknoglobals)

Global variables are an input to functions that is not visible in the functions signature, complicate testing, reduces readability and increase the complexity of code. However, sometimes global varaibles make sense. This tool skips such common scenarios. This tool can be used in CI, albeit it is very strict. This tool is useful for investigations. — [@leighmcculloch](https://github.com/leighmcculloch)


```
gochecknoglobals ./...
```

Example
```
/Users/nikolaydubina/Workspace/hugo/common/paths/path.go:64:5: fpb is a global variable
/Users/nikolaydubina/Workspace/hugo/common/paths/url.go:50:5: pb is a global variable
/Users/nikolaydubina/Workspace/hugo/common/text/position.go:52:5: positionStringFormatfunc is a global variable
/Users/nikolaydubina/Workspace/hugo/common/text/transform.go:26:5: accentTransformerPool is a global variable
/Users/nikolaydubina/Workspace/hugo/common/herrors/error_locator.go:40:5: SimpleLineMatcher is a global variable
```

Requirements
```
go install 4d63.com/gochecknoglobals@latest
```

### [⏫](#contents)➡ Detect slices that could be preallocated with [prealloc](https://github.com/alexkohler/prealloc)

Preallocating slices can sometimes significantly improve performance. This tool detects common scenarions where preallocating can be beneficial. This tool is not using `golang.org/x/tools/go/analysis` toolchain. — [@alexkohler](https://github.com/alexkohler)


```
prealloc ./...
```

Example
```
tools/gopls/internal/lsp/source/completion/completion.go:1484 Consider preallocating paths
tools/gopls/internal/lsp/source/completion/package.go:54 Consider preallocating items
tools/gopls/internal/lsp/template/symbols.go:205 Consider preallocating ans
tools/gopls/internal/lsp/template/completion.go:199 Consider preallocating working
tools/gopls/internal/lsp/tests/util.go:32 Consider preallocating notePositions
tools/gopls/internal/lsp/tests/util.go:240 Consider preallocating paramParts
tools/gopls/internal/lsp/tests/util.go:282 Consider preallocating result
tools/gopls/internal/lsp/tests/util.go:309 Consider preallocating got
```

Requirements
```
go install github.com/alexkohler/prealloc@latest
```

### [⏫](#contents)➡ Detect unnecessary import aliases with [unimport](https://github.com/alexkohler/unimport)

It is common guideline to avoid renaming imports unless there are collisions. This tool detects where original pacakge name would not collide. This tool is useful for investigations. This tool is not using `golang.org/x/tools/go/analysis` toolchain. — [@alexkohler](https://github.com/alexkohler)


```
unimport ./...
```

Example
```
pkg/apis/apiserverinternal/v1alpha1/zz_generated.conversion.go:29 unnecessary import alias runtime
pkg/apis/apiserverinternal/v1alpha1/zz_generated.conversion.go:30 unnecessary import alias apiserverinternal
pkg/apis/apps/v1/zz_generated.conversion.go:25 unnecessary import alias unsafe
pkg/apis/apps/v1/zz_generated.conversion.go:30 unnecessary import alias conversion
pkg/apis/apps/v1/zz_generated.conversion.go:31 unnecessary import alias runtime
pkg/apis/apps/v1/zz_generated.conversion.go:32 unnecessary import alias intstr
pkg/apis/apps/v1/zz_generated.conversion.go:33 unnecessary import alias apps
pkg/apis/apps/v1/zz_generated.conversion.go:34 unnecessary import alias core
pkg/apis/apps/v1beta1/zz_generated.conversion.go:25 unnecessary import alias unsafe
pkg/apis/apps/v1beta1/zz_generated.conversion.go:27 unnecessary import alias v1beta1
pkg/apis/apps/v1beta1/zz_generated.conversion.go:30 unnecessary import alias conversion
pkg/apis/apps/v1beta1/zz_generated.conversion.go:31 unnecessary import alias runtime
```

Requirements
```
go install github.com/alexkohler/unimport@latest
```

### [⏫](#contents)➡ Detect naked returns with [nakedret](https://github.com/alexkohler/nakedret)

It is common guideline to avoid [naked returns](https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters). Naked return is when function has named return, and return statement does not specify value. This tool is useful for investigations. — [@alexkohler](https://github.com/alexkohler)


```
nakedret ./...
```

Example
```
/kubernetes/pkg/controller/podautoscaler/replica_calculator.go:421:2: naked return in func `groupPods` with 44 lines of code
/kubernetes/pkg/kubelet/container/helpers.go:374:2: naked return in func `MakePortMappings` with 36 lines of code
/kubernetes/pkg/kubelet/config/config.go:350:2: naked return in func `filterInvalidPods` with 17 lines of code
/kubernetes/pkg/kubelet/config/config.go:449:3: naked return in func `checkAndUpdatePod` with 38 lines of code
/kubernetes/pkg/kubelet/config/config.go:471:2: naked return in func `checkAndUpdatePod` with 38 lines of code
/kubernetes/cmd/kube-controller-manager/app/controllermanager.go:717:2: naked return in func `createClientBuilders` with 19 lines of code
/kubernetes/pkg/proxy/topology.go:77:3: naked return in func `CategorizeEndpoints` with 98 lines of code
/kubernetes/pkg/proxy/topology.go:111:3: naked return in func `CategorizeEndpoints` with 98 lines of code
/kubernetes/pkg/proxy/topology.go:119:3: naked return in func `CategorizeEndpoints` with 98 lines of code
/kubernetes/pkg/proxy/topology.go:137:2: naked return in func `CategorizeEndpoints` with 98 lines of code
```

Requirements
```
go install github.com/alexkohler/nakedret/cmd/nakedret@latest
```

### [⏫](#contents)➡ Detect mixing pointer and value method receivers with [smrcptr](https://github.com/nikolaydubina/smrcptr)

Mixing pointer and value method receivers for the same type is discouraged, as per commong guideline [Go wiki](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type) and [Google Go style guide](https://google.github.io/styleguide/go/decisions#receiver-type). — [@nikolaydubina](https://github.com/nikolaydubina)


```
smrcptr ./...
```

```go
type Pancake struct{}

func NewPancake() Pancake { return Pancake{} }

func (s *Pancake) Fry() {}

func (s Pancake) Bake() {}
```

Example
```
smrcptr/internal/bakery/pancake.go:7:1: Pancake.Fry uses pointer
smrcptr/internal/bakery/pancake.go:9:1: Pancake.Bake uses value
```

Requirements
```
go install github.com/nikolaydubina/smrcptr@latest
```

### [⏫](#contents)➡ Detect vertical function ordering with [vertfn](https://github.com/nikolaydubina/vertfn)

Vertical function ordering is declaring functions before they are used. Based on 'Clean Code' by Robert.C.Martin. — [@nikolaydubina](https://github.com/nikolaydubina)


```
vertfn --verbose ./...
```

<div align="center"><img src="https://github.com/nikolaydubina/vertfn/blob/master/doc/code-dep-viz.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikolaydubina/vertfn@latest
```

### [⏫](#contents)➡ Calculate Cognitive Complexity with [gocognit](https://github.com/uudashr/gocognit)

Congitive Complexity as defined in this tool can be more illustrative than Cyclometric Complexity. Research paper ["Cognitive Complexity - a new way of measuring understandability"](https://www.sonarsource.com/docs/CognitiveComplexity.pdf), 2021. — [@uudashr](https://github.com/uudashr)


```
gocognit .
```

```go
// Complexity Cyclomatic=4 Cognitive=7
// Cognitive complexity give higher score compare to cyclomatic complexity.
func SumOfPrimes(max int) int {         // +1
    var total int
    for i := 1; i < max; i++ {          // +1 (cognitive +1, nesting)
        for j := 2; j < i; j++ {        // +1 (cognitive +2, nesting)
            if i%j == 0 {               // +1
                continue OUT
            }
        }
        total += i
    }
    return total
}

// Complexity Cyclomatic=4 Cognitive=1
// Cognitive complexity give lower score compare to cyclomatic complexity.
func GetWords(number int) string {      // +1
    switch number {
        case 1:                         // +1 (cognitive 0)
            return "one"
        case 2:                         // +1 (cognitive 0)
            return "a couple"
        case 3:                         // +1 (cognitive 0)
            return "a few"
        default:
            return "lots"
    }
}
```

Example
```
21 main (BasicSymtabConverter).SymtabFileToTreemap basic_converter.go:23:1
12 symtab parseGoSymtabLine symtab/go_symtab_parser.go:37:1
11 main main main.go:30:1
8 symtab EqSymbolName symtab/symbol_name_parser.go:12:1
7 symtab ParseSymbolName symtab/symbol_name_parser.go:32:1
7 symtab Test_parseGoSymtabLine symtab/go_symtab_parser_private_test.go:5:1
4 symtab Test_ParseSymbolName symtab/symbol_name_parser_private_test.go:5:1
3 main updateNodeNamesWithByteSize main.go:99:1
3 main unique basic_converter.go:119:1
3 symtab (GoSymtabParser).ParseSymtab symtab/go_symtab_parser.go:14:1
2 fmtbytecount ByteCountIEC fmtbytecount/format_bytecount.go:3:1
```

Requirements
```
go install github.com/uudashr/gocognit/cmd/gocognit@latest
```

### [⏫](#contents)➡ Calculate Cyclomatic Complexity with [gocyclo](https://github.com/fzipp/gocyclo)

Cyclomatic complexity is a code quality metric which can be used to identify code that needs refactoring. It measures the number of linearly independent paths through a function's source code. For example, excessive usage of nested `if` and `for` leads to increased cyclomatic complexity. This tool can report `top-N` and `over`, which makes it suitable for CI as a linter and manual investigation. — [@fzipp](https://github.com/fzipp)


```
gocyclo .
```

Example
```
$ gocyclo -over=5 .
34 examplemodule (*With32FieldsFeatureTransformer).Fit cmd/generate/tests/with32fieldsfp.go:48:1
24 main parseCode cmd/generate/parser.go:83:1
13 examplemodule (*AllTransformersFeatureTransformer).Fit cmd/generate/tests/alltransformersfp.go:27:1
12 examplemodule (*EmployeeFeatureTransformer).Fit cmd/generate/tests/employeefp.go:26:1
11 transformers (*CountVectorizer).TransformInplace transformers/textprocesors.go:84:1
11 structtransformer (*StructTransformer).Transform structtransformer/structtransformer.go:38:1
11 examplemodule (*LargeMemoryTransformerFeatureTransformer).Fit cmd/generate/tests/largememorytransformerfp.go:25:1
10 examplemodule (*WeirdTagsFeatureTransformer).Fit cmd/generate/tests/weirdtagsfp.go:24:1
8 transformers (*SampleNormalizerL2).TransformInplace transformers/samplenormalizers.go:58:1
```

Requirements
```
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

### [⏫](#contents)➡ Calculate age of comments with [go-commentage](https://github.com/nikolaydubina/go-commentage)

This go vet compatible tool analyses AST and git and collects details on how far comments drift from code they describe. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go-commentage -min-days-behind 360 ./...
```

Example
```
kubernetes/pkg/util/ipset/ipset.go:283:1: "CreateSet": doc_last_updated_behind_days(1336.83)
kubernetes/pkg/util/ipset/ipset.go:296:1: "createSet": doc_last_updated_behind_days(1603.17)
kubernetes/pkg/util/ipset/ipset.go:320:1: "AddEntry": doc_last_updated_behind_days(1578.10)
kubernetes/pkg/util/ipset/ipset.go:332:1: "DelEntry": doc_last_updated_behind_days(1578.10)
kubernetes/pkg/util/ipset/ipset.go:340:1: "TestEntry": doc_last_updated_behind_days(450.07)
```

Requirements
```
# get latest version of git
go install github.com/nikolaydubina/go-commentage@latest
```

### [⏫](#contents)➡ (archived) Ensure `if` statements using short assignment with [ifshort](https://github.com/esimonov/ifshort)

Linter for checking that your code uses short syntax for `if` statements whenever possible. However, as of `2023-05-26`, it is not maitaned and is not working. — [@esimonov](https://github.com/esimonov)


```
ifshort ./...
```

```go
// bad
func someFunc(k string, m map[string]interface{}) {
  _, ok := m[k]
  if !ok {
    return
  }

  err := otherFunc1()
  if err != nil {
    otherFunc2(err)
  }
}

// good
func someFunc(k string, m map[string]interface{}) {
  if _, ok := m[k]; !ok {
    return
  }

  if err := otherFunc1(); err != nil {
    otherFunc2(err)
  }
}
```

Requirements
```
go install github.com/esimonov/ifshort@latest
```

### [⏫](#contents)➡ Perform Taint Analysis with [taint](https://github.com/picatz/taint)

Taint analysis is a technique for identifying the flow of sensitive data through a program. It can be used to identify potential security vulnerabilities, such as SQL injection or cross-site scripting (XSS) attacks, by understanding how this data is used and transformed as it flows through the code. This package provides tools to performs such analysis. Included tool is performing SQL injection taint analysis. — [@picatz](https://github.com/picatz)


```
sqli main.go
```

```go
package main

import (
        "database/sql"
        "net/http"
)

func business(db *sql.DB, q string) {
        db.Query(q) // potential sql injection
}

func run() {
        db, _ := sql.Open("sqlite3", ":memory:")

        mux := http.NewServeMux()

        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                business(db, r.URL.Query().Get("sql-query"))
        })

        http.ListenAndServe(":8080", mux)
}

func main() {
        run()
}
```

Example
```
./sql/injection/testdata/src/example/main.go:9:10: potential sql injection
```

Requirements
```
go install github.com/picatz/taint/cmd/sqli@latest
```

### [⏫](#contents)➡ Visualize struct layout with [structlayout](https://github.com/dominikh/go-tools/tree/master/cmd/structlayout)

Display the byte offset and size of each field, respecting alignment/padding. — [@dominikh](https://github.com/dominikh)


```
structlayout -json bytes Buffer | structlayout-svg -t "bytes.Buffer" > /tmp/struct.svg
```

<div align="center"><img src="https://github.com/dominikh/go-tools/blob/master/images/screenshots/struct.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/ajstarks/svgo/structlayout-svg@latest
go install honnef.co/go/tools/cmd/structlayout@latest
```

### [⏫](#contents)➡ Rely on compiler for stricter Enums

For compile time blocking of: accidental arithmetics; implicit cast of untyped constants; all operators except `==` and `!=`; — simply wrap into a struct in separate package and do not export field. [example](http://github.com/nikolaydubina/go-enum-example).

```go
package color

type Color struct{ c uint }

var (
  Undefined = Color{}
  Red       = Color{1}
  Green     = Color{2}
  Blue      = Color{3}
)
```


## Code Generation

### [⏫](#contents)➡ Run `go:generate` in parallel

Official Go team [encourages](https://github.com/golang/go/issues/20520) to run sequentially. However, in certain situations, such as lots of mocks, parallelization helps a lot, albeit, you should consider including your generated files in git. The solution bellow spawns multiple processes, each per pkg.


```
grep -rnw "go:generate" -E -l "${1:-*.go}" . | xargs -L1 dirname | sort -u | xargs -P 8 -I{} go generate {}
```


### [⏫](#contents)➡ Generate `String` method for enum types

This is an official tool for generating `String` for enums. It supports overrides via comments. — official Go team

```go
package painkiller

//go:generate stringer -type=Pill -linecomment

type Pill int

const (
  Placebo Pill = iota
  Ibuprofen
  Paracetamol
  PillAspirin   // Aspirin
  Acetaminophen = Paracetamol
)

// "Acetaminophen"
var s string = Acetaminophen.String()
```

Requirements
```
go install golang.org/x/tools/cmd/stringer@latest
```

### [⏫](#contents)➡ Modify struct field tags with [gomodifytags](https://github.com/fatih/gomodifytags)

This tool makes it easy to update, add or delete the tags and options in a struct field. You can add new tags, update existing tags (such as appending a new key, i.e: db, xml, etc..) or remove existing tags. It's intended to be used by an editor, but also has modes to run it from the terminal. — [@fatih](https://github.com/fatih)

<div align="center"><img src="https://user-images.githubusercontent.com/438920/32691304-a1c7e47c-c716-11e7-977c-f4d0f8c616be.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/fatih/gomodifytags@latest
```

### [⏫](#contents)➡ Generate Table Driven Tests with [gotests](https://github.com/cweill/gotests)

This tool generates basic test placeholder. It is included into official Go plugin in VSCode and other major code editors. — [@cweill](https://github.com/cweill)

<div align="center"><img src="https://github.com/cweill/GoTests-Sublime/raw/master/gotests.gif" style="margin: 8px; max-height: 640px;"></div>



## Refactoring

### [⏫](#contents)➡ Replace symbol with `gofmt`

I found this in announcement [notice](https://github.com/golang/go/commit/2580d0e08d5e9f979b943758d3c49877fb2324cb) of Go 1.18 for changes to `interface{}` to `any`. This can be useful for other refactorings too.


```
gofmt -w -r 'interface{} -> any' .
```


### [⏫](#contents)➡ Keep consistent ordering of imports with [gci](https://github.com/daixiang0/gci)

This tool splits all import blocks into different sections, now support five section types: standard (e.g. 'fmt'); custom; default; blank; dot. It will keep each section sorted and keep ordering of sections consistent. — [@daixiang0](https://github.com/daixiang0)


```
gci write -s standard -s default -s "prefix(github.com/daixiang0/gci)" main.go
```

```go
// before
package main
import (
  "golang.org/x/tools"
  
  "fmt"
  
  "github.com/daixiang0/gci"
)

// after
package main
import (
    "fmt"

    "golang.org/x/tools"

    "github.com/daixiang0/gci"
)
```

Requirements
```
go install github.com/daixiang0/gci@latest
```

### [⏫](#contents)➡ Keep consistent ordering of imports with [goimportx](https://github.com/anqiansong/goimportx/tree/main)

This tool groups and sorts imports within groups. It keeps consitent ordering of groups. Detection of groups may be not always accurate. — [@anqiansong](https://github.com/anqiansong)


```
goimportx --file /path/to/file.go --group "system,local,third"
```

```go
package main

import (
  "flag"
  "io"
  "log"
  "os"

  "github.com/nikolaydubina/mdpage/page"
  "github.com/nikolaydubina/mdpage/render"
  yaml "gopkg.in/yaml.v3"
) 
```

Requirements
```
go install github.com/anqiansong/goimportx@latest
```

## Errors

### [⏫](#contents)➡ Pretty print `panic` messages with [panicparse](https://github.com/maruel/panicparse)

Read `panic` messages easier. Need to redirect STDERR to this tool with `panic` stack traces. The tool has HTML output and does lots of deduplication and enhancements. Refer to examples in original repo. — [@maruel](https://github.com/maruel)


```
go test -v |& pp
```

<div align="center"><img src="https://raw.githubusercontent.com/wiki/maruel/panicparse/parse.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/maruel/panicparse/v2/cmd/pp@latest
```

## Build

### [⏫](#contents)➡ Show compiler optimization decisions on heap and inlining

Building with `-m` flag will show decisions of compiler on inlining and heap escape. This can help you to validate your understanding of your code and optimize it.


```
go build -gcflags="-m -m" . 2>&1 | grep inline
```

Example
```
...
./passengerfp.go:25:6: cannot inline (*PassengerFeatureTransformer).Fit: function too complex: cost 496 exceeds budget 80
...
./passengerfp.go:192:6: can inline (*PassengerFeatureTransformer).NumFeatures with cost 35 as: method(*PassengerFeatureTransformer) func() int { if e == nil { return 0 }; count := 6; count += (*transformers.OneHotEncoder).NumFeatures(e.Sex); count += (*transformers.OneHotEncoder).NumFeatures(e.Embarked); return count }
...
./passengerfp.go:238:43: inlining call to transformers.(*OneHotEncoder).FeatureNames
./passengerfp.go:238:43: inlining call to transformers.(*OneHotEncoder).NumFeatures
...
./passengerfp.go:151:7: parameter e leaks to {heap} with derefs=0:
./passengerfp.go:43:11: make(map[string]uint) escapes to heap
```


### [⏫](#contents)➡ Disable inlining

Usually you may not need it, but can reduce binary size and even improve performance.


```
go build -gcflags="-l" .
```


### [⏫](#contents)➡ Aggressive inlining

Usually you may not need it, but can improve performance. This includes mid-stack inlining.


```
go build -gcflags="-l -l -l -l" .
```


### [⏫](#contents)➡ Manually disable or enable `cgo`

Disable `cgo` with `CGO_ENABLED=0` and enable with `CGO_ENABLED=1`. If you don't, `cgo` may end-up being enabled or code dynamically linked if, for example, you use some `net` or `os` packages. You may want to disable `cgo` to improve performance, since complier and runtime would have easier job optimizing code. This also should reduce your image size, as you can have alpine image with less shared libraries.


### [⏫](#contents)➡ Include metadata in binary during compilation with `ldflags`

You can pass metadata through compiler to your binary. This is useful for including things like git commit, database schema version, integrity hashes. Variables can only be strings.


```
go build -v -ldflags="-X 'main.Version=v1.0.0'"
go build -v -ldflags="-X 'my/pkg/here.Variable=some-string'"
```

```go
package main

var Version string

func main() {
  // Version here has some value
  ...
}
```


### [⏫](#contents)➡ Make treemap breakdown of Go executable binary with [go-binsize-treemap](https://github.com/nikolaydubina/go-binsize-treemap)

Useful for studying Go compiler, large projects, projects with C/C++ and `cgo`, 3rd party dependencies, embedding. However, total size may not be something to worry about for your executable. — [@nikolaydubina](https://github.com/nikolaydubina)


```
go tool nm -size <binary finename> | go-binsize-treemap > binsize.svg
```

<div align="center"><img src="https://github.com/nikolaydubina/go-binsize-treemap/blob/main/docs/hugo.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikolaydubina/go-binsize-treemap@latest
```

### [⏫](#contents)➡ Custom import path

Go can automatically fetch from custom http/https servers using `<meta>` tag to discover how to fetch code. There are multiple tools that can help set this up. This can help for security and analytics. This is also known as vanity URLs. [documentation](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths).

```
# some notable examples
golang.org/x/exp
go.uber.org/multierr
honnef.co/go/tools/cmd/staticcheck
```


### [⏫](#contents)➡ Custom import path with [govanityurls](https://github.com/GoogleCloudPlatform/govanityurls)

Simple HTTP server that lets you host custom import paths for your Go packages. — Google


```
govanityurls
```

Requirements
```
# get custom domain
# define YAML spec for redirection
go install github.com/GoogleCloudPlatform/govanityurls@latest
```

### [⏫](#contents)➡ Custom import path with [sally](https://github.com/uber-go/sally)

Simple HTTP server that lets you host custom import paths for your Go packages. — Uber


```
sally
```

Requirements
```
# get custom domain
# define YAML spec for redirection
go install go.uber.org/sally@latest
```

### [⏫](#contents)➡ Custom import path with [kkn.fi/vanity](https://kkn.fi/vanity)

Simple HTTP server that lets you host custom import paths for your Go packages. — [@kare](https://github.com/kare)


```
vanity
```

Requirements
```
# get custom domain
go get kkn.fi/vanity
```

### [⏫](#contents)➡ Custom import path enforcement

When import path is using custom domain, it is possible to block code from compilation unless it is used. This can help ensure security and prevent breaking changes. [documentation](https://pkg.go.dev/cmd/go#hdr-Import_path_checking).

```go
package pdf // import "rsc.io/pdf"
```


## Assembly

### [⏫](#contents)➡ Get assembly of Go code snippets online

Use [godbolt.org](https://godbolt.org) to compile and see assembly of short Go code. You can check different platforms and compilers including `cgo`. This tool is commonly used by C++ community. — [@mattgodbolt](https://github.com/mattgodbolt)

<div align="center"><img src="./img/godbolt.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Get Go SSA intermediary representation with [ssaplayground](https://github.com/golang-design/ssaplayground)

Check what does Go compiler do. Might be useful if you trying to optimize some code or learn more about compiler. https://golang.design/gossa. — [@changkun](https://github.com/changkun)

<div align="center"><img src="https://github.com/golang-design/ssaplayground/blob/main/public/assets/screen.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ View Go assembly interactively with [lensm](https://github.com/loov/lensm)

Understand how Go is compiled better. — [@egonelbre](https://github.com/egonelbre)

<div align="center"><img src="https://github.com/loov/lensm/raw/main/screenshot.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install loov.dev/lensm@main
```

### [⏫](#contents)➡ Generate Go assembly in Go with [avo](https://github.com/mmcloughlin/avo)

Write better quality Go assembly quicker in Go language itself. This tool conveniently generates stub for Go code to call your generated assembly. Used by Go core. — [@mmcloughlin](https://github.com/mmcloughlin)

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


### [⏫](#contents)➡ Generate AST for code snippets

Access Go core AST mechanism to generate AST.

```go
package main

import (
  "go/ast"
  "go/parser"
  "go/token"
)

func main() {
  fs := token.NewFileSet()
  tr, _ := parser.ParseExpr("(3-1) * 5")
  ast.Print(fs, tr)
}
```

Example
```
0  *ast.BinaryExpr {
1  .  X: *ast.ParenExpr {
2  .  .  Lparen: -
3  .  .  X: *ast.BinaryExpr {
4  .  .  .  X: *ast.BasicLit {
5  .  .  .  .  ValuePos: -
6  .  .  .  .  Kind: INT
7  .  .  .  .  Value: "3"
8  .  .  .  }
9  .  .  .  OpPos: -
10  .  .  .  Op: -
11  .  .  .  Y: *ast.BasicLit {
12  .  .  .  .  ValuePos: -
13  .  .  .  .  Kind: INT
14  .  .  .  .  Value: "1"
15  .  .  .  }
16  .  .  }
17  .  .  Rparen: -
18  .  }
19  .  OpPos: -
20  .  Op: *
21  .  Y: *ast.BasicLit {
22  .  .  ValuePos: -
23  .  .  Kind: INT
24  .  .  Value: "5"
25  .  }
26  }
```


### [⏫](#contents)➡ Visualize Go SSA function using Graphviz with [go-ssaviz](https://github.com/SilverRainZ/go-ssaviz)

This tool provides a visual overview of Go SSA function using Graphviz. This is especially useful in SSA-based static analysis. This tool generates an HTML page that is easy to navigate. [demo](https://silverrainz.me/go-ssaviz/). — [@SilverRainZ](https://github.com/SilverRainZ)


```
go-ssaviz ./...
```

<div align="center"><img src="./img/go-ssaviz.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
# get graphviz
go install github.com/SilverRainZ/go-ssaviz@latest
```

### [⏫](#contents)➡ (archived) Make graph of AST with [astgraph](https://github.com/xiazemin/ast_graph)

This tool visualizes AST as graph, which may be useful to navigate and undertand Go AST. This tool has not been maintaned for a while. — [@xiazemin](https://github.com/xiazemin)

<div align="center"><img src="https://github.com/xiazemin/ast_graph/raw/master/tree.svg" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
graphviz
```

## Execution

### [⏫](#contents)➡ Run interactive Go interpreter with [gomacro](https://github.com/cosmos72/gomacro)

This is interactive Go interpreter and debugger with REPL, Eval, generics and Lisp-like macros. You can run functions, import 3rd patry packages. Can be useful for learning and experimentation. Some nice features: autocomplete; constant expressions arithmetics. As of `2023-06-02`, issues with importing 3rd paty package are possible. — [@cosmos72](https://github.com/cosmos72)


```
gomacro
```

Example
```
$ gomacro
gomacro> import "fmt"
gomacro> fmt.Println("hello, world!")
hello, world!
14      // int
<nil>   // error
gomacro>
```

Requirements
```
go install github.com/cosmos72/gomacro@latest
```

### [⏫](#contents)➡ Run Go function in shell with [gorram](https://github.com/natefinch/gorram)

Run Go one-liners. This tool will print to STDOUT the return of a function call. — [@natefinch](https://github.com/natefinch)


```
cat README.md | gorram crypto/sha1 Sum
echo 12345 | gorram encoding/base64 StdEncoding.EncodeToString
gorram net/http Get https://google.com
```

Requirements
```
go install github.com/natefinch/gorram@latest
```

### [⏫](#contents)➡ Run simple fileserver

It takes one line to run HTTP file server in Go. Akin to famous oneliner in Python `python3 -m http.server` and `python -m SimpleHTTPServer`. Run this file as usually `go run <filename>`.

```go
package main

import "net/http"

func main() { http.ListenAndServe(":9000", http.FileServer(http.Dir("."))) }
```

<div align="center"><img src="./img/simple-fs.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Create 3D visualization of concurrency traces with [gotrace](https://github.com/divan/gotrace)

Fresh artistic perspective on coroutines execution. There is no advanced functions and it is hard to analyze production systems. However, it could be interesting for educational purposes. — [@divan](https://github.com/divan)

<div align="center"><img src="https://github.com/divan/gotrace/blob/master/images/demo.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/divan/gotrace@latest
patch Go compiler, available via Docker
more instructions in original repo
```

## Monitoring

### [⏫](#contents)➡ Monitor goroutines with [grmon](https://github.com/bcicen/grmon)

Command line monitoring for goroutines. — [@bcicen](https://github.com/bcicen)


```
grmon
```

<div align="center"><img src="https://camo.githubusercontent.com/ff8303d0b302fcfaf670846eb4168ac3e70522a8d739491d5509abc6ffb236b8/68747470733a2f2f627261646c65792e636f6465732f7374617469632f696d672f67726d6f6e2e676966" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
# start pprof server or grmon in your Go process
go install github.com/bcicen/grmon@latest
```

### [⏫](#contents)➡ Monitor Go processes with [gops](https://github.com/google/gops)

Monitoring memory of Go processes, forcing GC, getting version of Go of processes. — Google


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

### [⏫](#contents)➡ Visualise Go runtime metrics in browser with [statsviz](https://github.com/arl/statsviz)

This tool exposes HTTP endpoint with charts for Go runtime such as heap, objects, goroutines, GC pauses, scheduler. This is useful drop-in solution for visualization of Go runtime. — [@arl](https://github.com/arl)

<div align="center"><img src="https://raw.githubusercontent.com/arl/statsviz/readme-docs/window.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go get github.com/arl/statsviz@latest
```

### [⏫](#contents)➡ Auto-Instrument all functions with [go-instrument](https://github.com/nikolaydubina/go-instrument)

Automatically instrument all functions with Open Telemetry Spans by code generation. Inserts errors into Spans. — [@nikolaydubina](https://github.com/nikolaydubina)


```
find . -name "*.go" | xargs -I{} go-instrument -app my-service -w -filename {}
```

<div align="center"><img src="https://github.com/nikolaydubina/go-instrument/raw/master/docs/fib-error.png?raw=true" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/nikolaydubina/go-instrument@latest
```

### [⏫](#contents)➡ Auto-Instrument all functions with [otelinji](https://github.com/hedhyw/otelinji)

Automatically instrument all functions with Open Telemetry Spans by code generation. Inserts errors into Spans. Supports custom templates and can be used for Open Tracing or any custom insertions. — [@hedhyw](https://github.com/hedhyw)


```
otelinji -w -filename input_file.go
otelinji -filename input_file.go > input_file.go
find . -name "*.go" | grep -v "vendor/\|.git/\|_test.go" | xargs -n 1 -t otelinji -w -filename
```

<div align="center"><img src="https://github.com/hedhyw/otelinji/blob/main/assets/diff.png?raw=true" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/hedhyw/otelinji/cmd/otelinji@latest
```

## Benchmarking

### [⏫](#contents)➡ Run benchmarks

Start here. This is the standard tool for benchmarking. It can also do advanced features like mutex profiles. More flags are in Go [documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags) and `go help testflag`.


```
go test -bench=. -benchmem -benchtime=10s ./...
```

Example
```
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/fpmoney
BenchmarkArithmetic/add_x1-10                     1000000000             0.5 ns/op           0 B/op           0 allocs/op
BenchmarkArithmetic/add_x100-10                     18430124            64.6 ns/op           0 B/op           0 allocs/op
BenchmarkJSONUnmarshal/small-10                      3531835           340.7 ns/op         198 B/op           3 allocs/op
BenchmarkJSONUnmarshal/large-10                      2791712           426.9 ns/op         216 B/op           3 allocs/op
BenchmarkJSONMarshal/small-10                        4379685           274.4 ns/op         144 B/op           4 allocs/op
BenchmarkJSONMarshal/large-10                        3321205           345.8 ns/op         192 B/op           5 allocs/op
PASS
ok      github.com/nikolaydubina/fpmoney    62.744s
```


### [⏫](#contents)➡ Table-driven benchmarks

Similar to tests, Go supports table-driven benchmarks, which is very helpful for fine gradation of meta-parameters. More details in the Go [blog](https://go.dev/blog/subtests).

```
func benchIteratorSelector(b *testing.B, n int) {
  // ... setup here
  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    err := myExpensiveFunc()
    if err != nil {
      b.Error(err)
    }
  }
}

func BenchmarkIteratorSelector(b *testing.B) {
  for _, q := range []int{100, 1000, 10000, 100000} {
    b.Run(fmt.Sprintf("n=%d", q), func(b *testing.B) {
      benchIteratorSelector(b, q)
    })
  }
}
```

Example
```
BenchmarkIteratorSelector/n=100-10    	  297792	      4265 ns/op	    5400 B/op	      13 allocs/op
BenchmarkIteratorSelector/n=1000-10   	   31400	     38182 ns/op	    9752 B/op	      16 allocs/op
BenchmarkIteratorSelector/n=10000-10  	    3134	    380777 ns/op	   89112 B/op	      24 allocs/op
BenchmarkIteratorSelector/n=100000-10 	     310	   3827292 ns/op	  912410 B/op	      32 allocs/op
```


### [⏫](#contents)➡ Generate benchmak CPU and Memory profiles

This is useful for identifying most time or memory consuming parts. Recommended to run for single benchmark at a time and with `-count` or `-benchtime` for better accuracy.


```
go test -bench=<my-benchmark-name> -cpuprofile cpu.out -memprofile mem.out ./...
```


### [⏫](#contents)➡ Visualize callgraph of profiles with `pprof`

Once you generate profiles, visualize them with `pprof`. Both memory and CPU profiles are supported. Many options are available. Refer to the link you get in SVG to how to interpret this graph. More official documentation [blog](https://go.dev/blog/pprof), [pkg-doc](https://pkg.go.dev/net/http/pprof). — official Go team


```
go tool pprof -svg cpu.out > cpu.svg
go tool pprof -svg mem.out > mem.svg
```

<div align="center"><img src="img/pprof_callgraph_cpu.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Visualize flamegraphs of profiles with `pprof`

Latest versions of `pprof` can also render [Flamegraphs](https://www.brendangregg.com/flamegraphs.html) for profiles. Make sure you set `-http` to start webserver. Then it is available in "View > Graph" in at http://0.0.0.0:80. — Google


```
pprof -http=0.0.0.0:80 cpu.out
```

<div align="center"><img src="img/pprof_flamegraph_cpu.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/google/pprof@latest
```

### [⏫](#contents)➡ Visualize profiles online

You can also visualize profiles with online tools are aloso available https://www.speedscope.app (cpu).

<div align="center"><img src="img/speedscope_cpu_profile.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Get delta between two benchmarks with [benchstat](https://golang.org/x/perf/cmd/benchstat)

This is standard way to compare two benchmark outputs. Names of benchmarks should be the same. Generate benchmarks as per usual. You would get multiple tables per dimension. If no output, then pass `-split="XYZ"`. If you do not see `delta`, then pass `-count=2` or more in benchmark generation. It is recommended to have alternative implementations in different packages, to keep benchmark names the same. — official Go team


```
benchstat -split="XYZ" old.txt new.txt
```

Example
```
name                    old time/op    new time/op    delta
JSONUnmarshal/small-10     502ns ± 0%     331ns ± 0%   -33.99%  (p=0.008 n=5+5)
JSONUnmarshal/large-10     572ns ± 0%     414ns ± 0%   -27.64%  (p=0.008 n=5+5)
JSONMarshal/small-10       189ns ± 0%     273ns ± 0%   +44.20%  (p=0.008 n=5+5)
JSONMarshal/large-10       176ns ± 0%     340ns ± 0%   +93.29%  (p=0.008 n=5+5)

name                    old alloc/op   new alloc/op   delta
JSONUnmarshal/small-10      271B ± 0%      198B ± 0%   -26.94%  (p=0.008 n=5+5)
JSONUnmarshal/large-10      312B ± 0%      216B ± 0%   -30.77%  (p=0.008 n=5+5)
JSONMarshal/small-10       66.0B ± 0%    144.0B ± 0%  +118.18%  (p=0.008 n=5+5)
JSONMarshal/large-10       72.0B ± 0%    192.0B ± 0%  +166.67%  (p=0.008 n=5+5)

name                    old allocs/op  new allocs/op  delta
JSONUnmarshal/small-10      6.00 ± 0%      3.00 ± 0%   -50.00%  (p=0.008 n=5+5)
JSONUnmarshal/large-10      7.00 ± 0%      3.00 ± 0%   -57.14%  (p=0.008 n=5+5)
JSONMarshal/small-10        2.00 ± 0%      4.00 ± 0%  +100.00%  (p=0.008 n=5+5)
JSONMarshal/large-10        2.00 ± 0%      5.00 ± 0%  +150.00%  (p=0.008 n=5+5)
```

Requirements
```
go install golang.org/x/perf/cmd/benchstat@latest
```

### [⏫](#contents)➡ Get summary of benchmarks with [benchstat](https://golang.org/x/perf/cmd/benchstat)

Compare multiple benchmarks. Names of benchmarks should be the same. Generate benchmarks as per usual. You would get multiple tables per dimension. If no output, then pass `-split="XYZ"`. It is recommended to have alternative implementations in different packages, to keep benchmark names the same. — official Go team


```
benchstat -split="XYZ" int.txt float32.txt fpmoney.txt
```

Example
```
name \ time/op          int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10  481ns ± 2%     502ns ± 0%     331ns ± 0%
JSONUnmarshal/large-10  530ns ± 1%     572ns ± 0%     414ns ± 0%
JSONMarshal/small-10    140ns ± 1%     189ns ± 0%     273ns ± 0%
JSONMarshal/large-10    145ns ± 0%     176ns ± 0%     340ns ± 0%

name \ alloc/op         int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10   269B ± 0%      271B ± 0%      198B ± 0%
JSONUnmarshal/large-10   288B ± 0%      312B ± 0%      216B ± 0%
JSONMarshal/small-10    57.0B ± 0%     66.0B ± 0%    144.0B ± 0%
JSONMarshal/large-10    72.0B ± 0%     72.0B ± 0%    192.0B ± 0%

name \ allocs/op        int.bench   float32.bench  fpmoney.bench
JSONUnmarshal/small-10   6.00 ± 0%      6.00 ± 0%      3.00 ± 0%
JSONUnmarshal/large-10   7.00 ± 0%      7.00 ± 0%      3.00 ± 0%
JSONMarshal/small-10     2.00 ± 0%      2.00 ± 0%      4.00 ± 0%
JSONMarshal/large-10     2.00 ± 0%      2.00 ± 0%      5.00 ± 0%
```

Requirements
```
go install golang.org/x/perf/cmd/benchstat@latest
```

### [⏫](#contents)➡ Continuous benchmarking

Track how benchmarks change in codebase over time. This is accomplished by running benchmarks for git commits, storing results, and visualizing difference. Running benchmarks can be in GitHub Actions or locally, storage can be in same repository `master` or dedicated branch, or standalone servers. It should be straightforward to setup this manually. Example of GitHub Action [spec](https://github.com/swaggest/rest/blob/master/.github/workflows/bench.yml) and [blog](https://dev.to/vearutop/continuous-benchmarking-with-go-and-github-actions-41ok) from [@vearutop](https://github.com/vearutop), and an example on how it produces a PR [comment](https://github.com/swaggest/rest/pull/88#issuecomment-1271540878).

<div align="center"><img src="img/cont-bench-vearutop.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Continuous benchmarking with [gobenchdata](https://github.com/bobheadxi/gobenchdata)

This tool uses `go test -bench` data in GitHub. It runs benchmarks, and uploads it as GitHub Pages for visualization. It is available as GitHub Action [gobenchdata](https://github.com/marketplace/actions/continuous-benchmarking-for-go). This is useful to see benchmark trends. — [@bobheadxi](https://github.com/bobheadxi)

<div align="center"><img src="https://github.com/bobheadxi/gobenchdata/raw/main/.static/demo-chart.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install go.bobheadxi.dev/gobenchdata@latest
```

### [⏫](#contents)➡ Continuous benchmarking with [benchdiff](https://github.com/willabides/benchdiff)

Automates comparing benchmarks with `benchstat` of two git references. It is available as GitHub Action [benchdiff](https://github.com/marketplace/actions/benchdiff) which runs `benchstat` of HEAD vs base branch. This is useful to see how benchmarks change with PRs in CI. — [@WillAbides](https://github.com/WillAbides)

<div align="center"><img src="img/cont-bench-willabides.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/willabides/benchdiff/cmd/benchdiff
```

### [⏫](#contents)➡ Continuous benchmarking with [cob](https://https://github.com/knqyf263/cob)

Automate comparing benchmarks with `benchstat` between `HEAD` and `HEAD^1`. It can be used to block CI pipelines if benchmarks deteriorate. It reports output as text in CLI. This cane be useful in CI or in local development. — [@knqyf263](https://github.com/knqyf263)

<div align="center"><img src="https://github.com/knqyf263/cob/raw/master/img/usage.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/knqyf263/cob@latest
```

### [⏫](#contents)➡ Generate live traces with `net/http/trace`

This will add endpoints to your your server. If you don't have server running already in your process, you can start one. Then you can point `pprof` tool to this data. For production, hide this endpoint in separate port and path. More details in documentation [trace](https://pkg.go.dev/cmd/trace), [net/http/pprof](https://pkg.go.dev/net/http/pprof).

```
package main

import (
  "log"
  "net/http"
  "net/http/pprof"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/custom_debug_path/profile", pprof.Profile)
  log.Fatal(http.ListenAndServe(":7777", mux))
}
```

Example
```
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
```


### [⏫](#contents)➡ Generate traces with `go test`

Produce a trace of execution of tests in pacakge.


```
go test -trace trace.out .
```


### [⏫](#contents)➡ View traces with `go tool trace`

You can view traces interactively in browser with standard Go tooling. This web tool also shows network blocking profile, synchronization blocking profile, syscall blocking profile, scheduler latency profile.


```
go tool trace trace.out
```

<div align="center"><img src="img/go_tool_trace_web.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Get wallclock traces with [fgtrace](https://github.com/felixge/fgtrace)

This tool can be more illustrative of Go traces than standard Go traces. — [@felixge](https://github.com/felixge)

```go
package main

import (
  "net/http"

  "github.com/felixge/fgtrace"
)

func main() {
  http.DefaultServeMux.Handle("/debug/fgtrace", fgtrace.Config{})
  http.ListenAndServe(":1234", nil)
}
```

<div align="center"><img src="https://github.com/felixge/fgtrace/raw/main/assets/fgtrace-example.png" style="margin: 8px; max-height: 640px;"></div>



### [⏫](#contents)➡ Get on/off CPU profiles with [fgprof](https://github.com/felixge/fgprof)

This tool can be more illustrative of Go profiles than standard Go profiling. — [@felixge](https://github.com/felixge)

```go
package main

import (
  "log"
  "net/http"
  _ "net/http/pprof"

  "github.com/felixge/fgprof"
)

func main() {
  http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
  go func() {
    log.Println(http.ListenAndServe(":6060", nil))
  }()

  // <code to profile>
}
```

<div align="center"><img src="https://github.com/felixge/fgprof/raw/master/assets/fgprof_pprof.png" style="margin: 8px; max-height: 640px;"></div>



## Documentation

### [⏫](#contents)➡ Make alternative documentation with [golds](https://github.com/go101/golds)

It has additional information like implementations of interface; promoted methods. The tool has nice minimalistic aesthetics. — [Tapir Liu](https://www.tapirgames.com)


```
golds ./...
```

<div align="center"><img src="img/golds.png" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install go101.org/golds@latest
```

### [⏫](#contents)➡ Read Go binary documentation in `man` format with [goman](https://github.com/appliedgocode/goman)

This tool fetches the repo's readme as a man page replacement. — [@christophberger](https://github.com/christophberger)


```
goman <mypackage>
```

<div align="center"><img src="https://github.com/appliedgocode/goman/raw/master/goman.gif" style="margin: 8px; max-height: 640px;"></div>


Requirements
```
go install github.com/appliedgocode/goman@lates
```

## Style Guide

- [Google](https://google.github.io/styleguide/go)

- [Uber](https://github.com/uber-go/guide)

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)


