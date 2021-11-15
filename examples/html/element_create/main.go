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

	_, tags, tagStates := doc.Tags()

	fmt.Println(tags, tagStates)

	for _, tagName := range tags {
		newElem := doc.CreateElement(tagName)

		if newElem == nil {
			continue
		}

		if !tagStates[tagName] {
			textNode := doc.CreateTextNode("This is a \"" + tagName + "\" node")
			newElem.Node().InsertChild(textNode.Node())
		}

		body := doc.BodyElement()
		body.Element().Node().InsertChild(newElem.Node())
	}

	docNode := doc.DomInterfaceNode()
	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	doc.Destroy()
}
