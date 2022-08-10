package main

import (
	"html/template"
	"net/http"
)

const tmpl = ""

func Concat(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	// bad
	tmpl := "<html><body><h1>" + customerId + "</h1></body></html>"

	return template.URL(tmpl)
}

func Concat_ok(r *http.Request) template.HTML {
	// ok
	tmpl := "<html><body><h1> abc </h1></body></html>"

	return template.URL(tmpl)
}
