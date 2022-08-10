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

func OthersThatAreFine(r *http.Request) template.HTML {
	// ok
	a := template.HTMLAttr("<html><body><h1>Hello, world</h1></body></html>")
	// ok
	a := template.JS("<html><body><h1>Hello, world</h1></body></html>")
	// ok
	a := template.URL("<html><body><h1>Hello, world</h1></body></html>")
	// ok
	a := template.CSS("<html><body><h1>Hello, world</h1></body></html>")
	// ok
	a := template.Srcset("<html><body><h1>Hello, world</h1></body></html>")
}

func OthersThatAreNOTFine(r *http.Request, data string) template.HTML {
	// bad
	a := template.HTMLAttr(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", data))
	// bad
	a := template.JS(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", data))
	// bad
	a := template.URL(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", data))
	// bad
	a := template.CSS(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", data))
	// bad
	a := template.Srcset(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", data))
}

func Concat(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	tmpl := "<html><body><h1>" + customerId + "</h1></body></html>"

	// bad
	return template.HTML(tmpl)
}

func ConcatBranch(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	doIt, err := strconv.ParseBool(r.URL.Query().Get("do"))
	if err != nil {
		return template.HTML("")
	}
	var tmpl string
	if doIt {
		tmpl = "<html><body><h1>" + customerId + "</h1></body></html>"
	} else {
		tmpl = ""
	}

	// bad
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
	tmpl, err := fmt.Printf("<html><body><h1>%s</h1></body></html>", customerId)
	if err != nil {
		return template.HTML("")
	}
	// bad
	return template.HTML(tmpl)
}

func FormattedInline(r *http.Request) template.HTML {
	customerId := r.URL.Query().Get("id")
	// bad
	return template.HTML(fmt.Sprintf("<html><body><h1>%s</h1></body></html>", customerId))
}

func main() {}
