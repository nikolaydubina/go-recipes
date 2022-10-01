package main

import (
	"net/http"

	"github.com/felixge/fgtrace"
)

func main() {
	http.DefaultServeMux.Handle("/debug/fgtrace", fgtrace.Config{})
	http.ListenAndServe(":1234", nil)
}
