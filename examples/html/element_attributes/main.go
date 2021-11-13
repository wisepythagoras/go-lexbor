package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<div id=\"hello-world\">Hello, World!</div>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	docNode := doc.DomInterfaceNode()

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	collection := html.CreateDomCollection(doc, 16)

	if collection == nil {
		fmt.Println("Collection was nil")
		return
	}

	body := doc.BodyElement()
	element := body.Element()

	elements := collection.DomElementsByTagName("div", element)

	for _, el := range elements {
		id := el.Attribute("id")

		if len(id) > 0 {
			fmt.Println("Found id:", id)
		}
	}

	doc.Destroy()
}
