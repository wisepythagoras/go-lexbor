package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	body := doc.BodyElement()
	body.Element().SetInnerHTML("<div>Hello, World!</div>")

	docNode := doc.DomInterfaceNode()

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	doc.Destroy()
}
