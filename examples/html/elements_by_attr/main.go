package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<div><p>hehehe</p><p id=\"hello\">My target</p></div>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	body := doc.BodyElement()

	elements, _ := body.Element().ElementsByAttr("id", "hello")

	fmt.Println("The element:")
	html.Serialize(elements[0].Node())

	element, _ := doc.GetElementById("hello")

	fmt.Println("The element:")
	html.Serialize(element.Node())

	doc.Destroy()
}
