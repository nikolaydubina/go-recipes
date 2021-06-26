<h1 align="center">Go recipes 🦩 </h1>
<p align="center">Handy commands to run in Go projects</p>

### Get graph of packages

> Use to find unexpected dependencies, visualize project. Works best for small number of packages.

```bash
$ go list -deps -json ./... | jq -c "select(.Standard!="true") | {from: .ImportPath, to: .Imports[]}" | jsonl-graph | dot -Tsvg > package-graph.svg
```
![package-graph](./docs/pacages-graph.svg)

### Scrape details about module graph

> Use to find low quality, unmaintained dependencies

```bash
$ go mod graph | import-graph -i=gomod | jsonl-graph -color-scheme=file://$PWD/basic.json | dot -Tsvg > output.svg
```
![gin-mod-graph-collected](./docs/gin-mod-graph-collected.svg)

#### Prerequisites

```bash
# get https://graphviz.org/download/
$ go install github.com/nikolaydubina/jsonl-graph@latest
$ go install github.com/nikolaydubina/import-graph@latest
```