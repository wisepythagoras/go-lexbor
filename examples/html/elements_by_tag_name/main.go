package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func getElementsByTagName(tagName string, elements *[]*html.Element, el *html.Element) {
	collection, _ := el.ElementsByTagName(tagName)
	localElements := *elements

	for _, el := range collection.Elements() {
		getElementsByTagName("div", elements, el)
		localElements = append(localElements, el)
	}

	elements = &localElements
}

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<div><p>hehehe</p><div id=\"hello\"><div>My target</div></div></div>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	body := doc.BodyElement()

	collection, _ := body.Element().ElementsByTagName("div")
	elements := collection.Elements()

	for _, el := range elements {
		getElementsByTagName("div", &elements, el)
	}

	fmt.Println("The elements:")
	for _, el := range elements {
		html.Serialize(el.Node())
		fmt.Println("----------")
	}

	fmt.Println("Children:", len(elements[0].Node().Children()))

	doc.Destroy()
}
