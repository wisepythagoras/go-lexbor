package html

// #include <lexbor/html/html.h>
import "C"
import "unsafe"

type Node struct {
	lexborNode *C.lxb_dom_node_t
}

func (n *Node) InsertChild(node *Node) {
	C.lxb_dom_node_insert_child(n.Ptr(), node.Ptr())
}

func (n *Node) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(n.Ptr()))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (n *Node) Ptr() *C.lxb_dom_node_t {
	return n.lexborNode
}
