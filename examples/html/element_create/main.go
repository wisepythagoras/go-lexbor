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

	docNode := doc.DomInterfaceNode()
	newElem := doc.CreateElement("span")
	textNode := doc.CreateTextNode("This is some text")
	newElem.Node().InsertChild(textNode.Node())

	body := doc.BodyElement()
	body.Element().Node().InsertChild(newElem.Node())

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	doc.Destroy()
}
