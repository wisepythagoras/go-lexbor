package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type DomText struct {
	lexborDomText *C.lxb_dom_text_t
}

func (d *DomText) Node() *Node {
	lxbNode := (*C.lxb_dom_node_t)(unsafe.Pointer(d.Ptr()))
	return &Node{ptr: lxbNode}
}

func (d *DomText) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(d.Ptr()))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (d *DomText) Ptr() *C.lxb_dom_text_t {
	return d.lexborDomText
}
