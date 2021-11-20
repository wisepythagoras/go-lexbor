package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<div><p class=\"a b c\">hehehe</p><p class=\"hello abc\">My target</p></div>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	body := doc.BodyElement()

	collection, _ := body.Element().ElementsByClassName("hello")

	fmt.Println("The element:")
	html.Serialize(collection.Elements()[0].Node())

	collection, _ = body.Element().ElementsByClassName("a")

	fmt.Println("The element:")
	html.Serialize(collection.Elements()[0].Node())

	doc.Destroy()
}
