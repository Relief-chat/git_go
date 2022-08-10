package main

import "fmt"
import "html/template"

func main() {
    var g = "foo"

    // bad
    const a template.HTML = fmt.Sprintf("<a href=%q>link</a>")
    // bad
    var b template.CSS = "a { text-decoration: underline; } "

    // bad
    var c template.HTMLAttr =  fmt.Sprintf("herf=%q")

    // bad
    const d template.JS = "{foo: 'bar'}"

    // bad
    var e template.JSStr = "setTimeout('alert()')";

    // bad
    var f template.Srcset = g;

    // ok 
    tmpl, err := template.New("test").ParseFiles("file.txt")


    myTpl.Execute(w, a);
}
