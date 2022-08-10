package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Fine(r *http.Request) template.HTML {
	// ok
	return template.HTML("<html><body><h1>Hello, world</h1></body></html>")
}

func AlsoFine(r *http.Request) template.HTML {
	// ok
	return template.HTML("<html><body><h1>" + "Hello, world</h1></body></html>")
}

func Concat(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	// bad
	tmpl := "<html><body><h1>" + customerId + "</h1></body></html>"

	return template.HTML(tmpl)
}

func ConcatInline(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")

	// bad
	return template.HTML("<html><body><h1>" + customerId + "</h1></body></html>")
}

func ConcatInlineOneside(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")

	// bad
	return template.HTML("<html><body><h1>" + customerId)
}

func Formatted(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	// bad
	tmpl, err := fmt.Printf("<html><body><h1>%s</h1></body></html>", customerId)
	if err != nil {
		return template.HTML("")
	}
	return template.HTML(tmpl)
}

func FormattedInline(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	// bad
	return template.HTML(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", customerId))
}

func main() {}
