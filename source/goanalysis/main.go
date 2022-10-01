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
