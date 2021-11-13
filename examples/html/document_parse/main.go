package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.Parse("<div><p>blah-blah-blah</div>")

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	docNode := doc.DomInterfaceNode()

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	doc.Destroy()
}
