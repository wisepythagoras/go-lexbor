package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<head><title>  Some   - title...   </title></head>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	docNode := doc.DomInterfaceNode()

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	fmt.Println("The title is:", doc.Title())

	if !doc.ChangeTitle("This is my new title") {
		fmt.Println("Failed to change title")
		return
	}

	html.Serialize(docNode)

	doc.Destroy()
}
