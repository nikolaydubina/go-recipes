<h1 align="center">Go Recipes 🦩 </h1>
<p align="center">Handy well-known and less-known tools for Go projects</p>

## Contents

- Tests
  + [➡ Make treemap of code coverage](#-make-treemap-of-code-coverage)
  + [➡ Get packages without tests](#-get-packages-without-tests)
  + [➡ Browse code coverage by file](#-browse-code-coverage-by-file)
  + [➡ Make histogram of Go files per package](#-make-histogram-of-go-files-per-package)
  + [➡ Run tests sequentially](#-run-tests-sequentially)
  + [➡ Run tests in parallel](#-run-tests-in-parallel)
- Dependencies
  + [➡ Get Go version of current module](#-get-go-version-of-current-module)
  + [➡ Get Go versions of upstream modules](#-get-go-versions-of-upstream-modules)
  + [➡ Get directly dependent modules that can be upgraded](#-get-directly-dependent-modules-that-can-be-upgraded)
  + [➡ Get upstream modules without Go version](#-get-upstream-modules-without-go-version)
  + [➡ Get available module versions](#-get-available-module-versions)
  + [➡ Make graph of upstream packages](#-make-graph-of-upstream-packages)
  + [➡ Scrape details about upstream modules and make graph](#-scrape-details-about-upstream-modules-and-make-graph)
  + [➡ Scrape licences of upstream dependencies](#-scrape-licences-of-upstream-dependencies)
  + [➡ Explore upstream dependencies interactively](#-explore-upstream-dependencies-interactively)
  + [➡ Make graph of function calls in package](#-make-graph-of-function-calls-in-package)
- Assembley
  + [➡ Get assembly of Go code snippets online](#-get-assembly-of-go-code-snippets-online)
  + [➡ Get Go compiler SSA intermediary representation](#-get-go-compiler-ssa-intermediary-representation)
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

## Tests

### ➡ Make treemap of code coverage

Visualize distribution of code coverage in your project.
This helps to identify code areas with high and low coverage.
Useful when you have large project with lots of files and packages.
This 2D image-hash of your project should be more representative than a single number.

First make profile with
```
go test -coverprofile cover.out ./...
```

Then turn coverprofile into SVG
```
go-cover-treemap -coverprofile cover.out > out.svg
```

<div align="center">
<img src="./img/hugo-code-coverage.svg" style="margin: 8px; max-height: 640px;">
</div>


<details><summary>Requirements</summary>

```
go install github.com/nikolaydubina/go-cover-treemap@latest
```

</details>
  
---

### ➡ Get packages without tests

If code coverage does not report packages without tests.
This should be fast for CI.

```
go list -json ./... | jq -rc 'select((.TestGoFiles | length)==0) | .ImportPath'
```

Example
```
github.com/gin-gonic/gin/ginS
github.com/gin-gonic/gin/internal/json
```

<details><summary>Requirements</summary>

```
https://stedolan.github.io/jq/download/
```

</details>
  
---

### ➡ Browse code coverage by file

This is very helpful tool from the official Go toolchain.

First make profile
```
go test -coverprofile cover.out ./...
```

Then open in browser
```
go tool cover -html=cover.out
```

<div align="center">
<img src="./img/tool-cover-html.png" style="margin: 8px; max-height: 640px;">
</div>

---

### ➡ Make histogram of Go files per package

Use this to see when package is too big or too small.
Adjust histogram length to maximum value.

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

<details><summary>Requirements</summary>

```
https://stedolan.github.io/jq/download/
```
  
</details>

---

### ➡ Run tests sequentially

This is in cases when you need to synchronize tests, for example in integration tests that share environment.
[Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).

```
go test -p 1 -parallel 1 ./...
```

---

### ➡ Run tests in parallel

Add `t.Parallel` to your tests case function bodies.
As per documentation, by default `-p=GOMAXPROCS` and `-parallel=GOMAXPROCS` when you run `go test`.
Different packages by default run in parallel, and tests within package can be enforced to run in parallel too.
Make sure to copy test case data to new variable, why explained [here](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721).
[Official documentation](https://pkg.go.dev/cmd/go#hdr-Testing_flags).

Example
```go
    ...
    for _, tc := range tests {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            ...
```

## Dependencies

### ➡ Get Go version of current module

Use this in CI to setup correct Go version automatically from `go.mod`.

```
go mod edit -json | jq -r .Go
```

Example
```
1.16
```

<details><summary>Requirements</summary>

```
https://stedolan.github.io/jq/download/
```

</details>

---

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

<details><summary>Requirements</summary>
  
```
https://stedolan.github.io/jq/download/
```

</details>

---

### ➡ Get directly dependent modules that can be upgraded

Use this to keep your modules updated.
Similar function is integrated in VSCode official Go plugin and GoLand.

```
go list -u -m $(go list -m -f '{{.Indirect}} {{.}}' all | grep '^false' | cut -d ' ' -f2) | grep '\['
```

Example
```
github.com/goccy/go-json v0.5.1 [v0.7.3]
github.com/golang/protobuf v1.3.3 [v1.5.2]
github.com/json-iterator/go v1.1.9 [v1.1.11]
```
---

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

<details><summary>Requirements</summary>

```
https://stedolan.github.io/jq/download/
```
  
</details>

---

### ➡ Get available module versions

This works even if you did not download or install module locally.
This is useful to check to which version you can upgrade to, what is the latest version, and whether there are v2+ major versions recognized by Go toolchain.

```
go list -m -versions github.com/google/gofuzz
```

Example
```
github.com/google/gofuzz v1.0.0 v1.1.0 v1.2.0
```
---

### ➡ Make graph of upstream packages

Use to find unexpected dependencies or visualize project.
Works best for small number of packages, for large projects use `grep` to narrow down subgraph.
Without `-deps` only for current module.

```
go list -deps -json ./... | jq -c 'select(.Standard!=true) | {from: .ImportPath, to: .Imports[]}' | jsonl-graph | dot -Tsvg > package-graph.svg
```

<div align="center">
<img src="./img/packages-graph.svg" style="margin: 8px; height: 640px;">
</div>

<details><summary>Requirements</summary>

```
https://stedolan.github.io/jq/download/
https://graphviz.org/download/
$ go install github.com/nikolaydubina/import-graph@latest
$ go install github.com/nikolaydubina/jsonl-graph@latest
```
  
</details>

---

### ➡ Scrape details about upstream modules and make graph

Use to find low quality or unmaintained dependencies.

```
go mod graph | import-graph -i=gomod | jsonl-graph -color-scheme=file://$PWD/basic.json | dot -Tsvg > output.svg
```

<div align="center">
<img src="./img/gin-mod-graph-collected.svg" style="margin: 8px; height: 640px;">
</div>

<details><summary>Requirements</summary>
  
```
https://graphviz.org/download/
$ go install github.com/nikolaydubina/import-graph@latest
$ go install github.com/nikolaydubina/jsonl-graph@latest
```
 
</details>

---

### ➡ Scrape licences of upstream dependencies

This is tool from Google. Might be useful to collect all the licences or checking if you can use the project for example in propriatary or commercial environment.

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

<details><summary>Requirements</summary>
  
```
go install github.com/google/go-licenses@latest
```
 
</details>

---

### ➡ Explore upstream dependencies interactively

This is a tool from one of creators of Go.
This tool should help explore dependencies and assist large refactorings.
Source: https://github.com/adonovan/spaghetti

<div align="center">
<img src="https://github.com/adonovan/spaghetti/blob/main/screenshot.png" style="margin: 8px; height: 640px;">
</div>

<details><summary>Requirements</summary>
  
```
go install github.com/adonovan/spaghetti@latest
```
 
</details>

---

### ➡ Make graph of function calls in package

This can be helpful to quickly track which packages current package is calling and why.

```
go-callvis .
```

<div align="center">
<img src="https://raw.githubusercontent.com/ofabry/go-callvis/master/images/syncthing.png" style="margin: 8px; max-height: 640px;">
</div>

<details><summary>Requirements</summary>
  
```
go install github.com/ofabry/go-callvis
```
 
</details>

## Assembly

### ➡ Get assembly of Go code snippets online

Use [godbolt.org](https://godbolt.org) to compile and see assembly of short Go code.
You can check different platforms and compilers including `cgo`.
This tool is commonly used by C++ community.

<div align="center">
<img src="./img/godbolt.png" style="margin: 8px; max-height: 640px;">
</div>

### ➡ Get Go compiler SSA intermediary representation

This tool allows to check what does Go compiler do.
Might be useful if you trying to optimize some code or learn more about compiler.
Website: https://golang.design/gossa 
Source: https://github.com/golang-design/ssaplayground

<div align="center">
<img src="https://github.com/golang-design/ssaplayground/blob/main/public/assets/screen.png" style="margin: 8px; height: 640px;">
</div>

## Execute

### ➡ Run Go function in shell

This is short and convenient for Go oneliners.
This tool will print to stdout the return of a function call.

Example
```
$ cat README.md | gorram crypto/sha1 Sum
c8faff3af2e6816800a8b83af8e3535872ec6120
```

Selected Usecases
- `echo 12345 | gorram encoding/base64 StdEncoding.EncodeToString`
- `cat myfile.txt | gorram crypto/sha1 Sum`
- `gorram net/http Get https://google.com`

<details><summary>Requirements</summary>
  
```
go install github.com/natefinch/gorram@latest
```
 
</details>

---

### ➡ Run simple fileserver

This is similar to famous oneliner in Python `python3 -m http.server` and `python -m SimpleHTTPServer`.

Create file like this
```go
package main

import (
    "net/http"
)

func main() {
    http.ListenAndServe(":9000", http.FileServer(http.Dir(".")))
}
```

Then run this file (e.g. if named `fs.go`)
```
go run fs.go
```

<div align="center">
<img src="./img/simple-fs.png" style="margin: 8px; max-height: 640px;">
</div>

--- 

### ➡ Monitor Go processes

This tool from Google has lots of useful features like monitoring memory of Go processes, forcing GC, getting version of Go of process.

```
$ gops
983   980    uplink-soecks  go1.9   /usr/local/bin/uplink-soecks
52697 52695  gops           go1.10  /Users/jbd/bin/gops
4132  4130   foops        * go1.9   /Users/jbd/bin/foops
51130 51128  gocode         go1.9.2 /Users/jbd/bin/gocode
```

<details><summary>Requirements</summary>
  
```
go install github.com/google/gops@latest
```
 
</details>

---

### ➡ Create 3D visualization of concurrency traces

This tool creates 3D visualization of coroutines execution.
There is no advanced functions and it is hard to analyze production systems.
However, it could be insteresting for educational purposes.
Source: https://github.com/divan/gotrace

<div align="center">
<img src="https://github.com/divan/gotrace/blob/master/images/demo.png" style="margin: 8px; height: 640px;">
</div>

<details><summary>Requirements</summary>
 
You may need to run Docker container with that tool or install locally and patch Go compiler.
More instructions in the original repo.
  
```
go install github.com/divan/gotrace
```
 
</details>

## Refactoring

### ➡ Replace symbol

I found this in annoncement [notice](https://github.com/golang/go/commit/2580d0e08d5e9f979b943758d3c49877fb2324cb) of Go 1.18 for changes to `interface{}` to `any`.
This can be useful for other refactorings too. 

```
gofmt -w -r 'interface{} -> any' .
```

## Errors

### ➡ Pretty print `panic` messages

This tool will be useful for reading `panic` messages.
Need to redirect STDERR to this tool with `panic` stack traces.
The tool has HTML outpout and does lots of deduplication and enhansements.
Refer to examples in original repo.

```
go test -v |& pp
```

<div align="center">
<img src="https://raw.githubusercontent.com/wiki/maruel/panicparse/parse.gif" style="margin: 8px; max-height: 640px;">
</div>

<details><summary>Requirements</summary>
  
```
go install github.com/maruel/panicparse/v2/cmd/pp@latest
```
 
</details>

## Build

### ➡ Manually disable or enable `cgo`

Disable `cgo` with `CGO_ENABLED=0` and enable with `CGO_ENABLED=1`.
If you don't, `cgo` may end-up being enabled or code dynamically linked if, for example, you use some `net` or `os` packages.
You may want to disable `cgo` to improve performance, since complier and runtime would have easier job optimizing code.
This also should reduce your image size, as you can have alpine image with less shared libraries.
