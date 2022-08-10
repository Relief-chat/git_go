import (
	"fmt"
	"github.com/lestrrat-go/libxml2/parser"
)

func vuln() {
	const s = "<!DOCTYPE d [<!ENTITY e SYSTEM \"file:///etc/passwd\">]><t>&e;</t>"
	// bad
	p := parser.New(parser.XMLParseNoEnt)
	doc, err := p.ParseString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Doc successfully parsed!")
	fmt.Println(doc)
}

func not_vuln() {
	const s = "<!DOCTYPE d [<!ENTITY e SYSTEM \"file:///etc/passwd\">]><t>&e;</t>"
	// ok 
	p := parser.New()
	doc, err := p.ParseString(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Doc successfully parsed!")
	fmt.Println(doc)
}