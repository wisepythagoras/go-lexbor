package main

import (
	"fmt"

	"github.com/wisepythagoras/go-lexbor/html"
)

var htmlChunks = []string{
	"<!DOCT",
	"YPE htm",
	"l>",
	"<html><head>",
	"<ti",
	"tle>HTML chun",
	"ks parsing</",
	"title>",
	"</head><bod",
	"y><div cla",
	"ss=",
	"\"bestof",
	"class",
	"\">",
	"good for me",
	"</div>",
}

func main() {
	doc := &html.Document{}
	doc.Create()
	success := doc.ParseChunks(htmlChunks)

	if !success {
		fmt.Println("Failed to create HTML Document")
		return
	}

	docNode := doc.DomInterfaceNode()

	fmt.Println("HTML Tree:")
	html.Serialize(docNode)

	doc.Destroy()
}
