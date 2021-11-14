package main

import (
	"fmt"
	"strconv"

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

	for i, el := range elements {
		if el.HasAttribute("id") {
			id := el.Attribute("id")

			fmt.Println("Found id:", id)
		}

		wasSet := el.SetAttribute("test-attr", "val-"+strconv.FormatInt(int64(i), 10))

		if !wasSet {
			fmt.Println("Unable to set test attribute on", i)
		}
	}

	element = collection.Element(0)
	attr := element.FirstAttribute()

	for attr != nil {
		fmt.Println(attr.QualifiedName(), "=>", attr.Value())
		attr = element.NextAttribute(attr)
	}

	domAttr := element.AttributeByName("id")

	if domAttr != nil && !domAttr.SetValue("alternate-id") {
		fmt.Println("Unable to set the new alternate id")
	}

	fmt.Println("New HTML Tree:")
	html.Serialize(docNode)

	collection.Destroy()
	doc.Destroy()
}
