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
