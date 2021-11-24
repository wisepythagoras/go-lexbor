package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type Node struct {
	ptr      *C.lxb_dom_node_t
	document *Document
}

func (n *Node) InsertChild(node *Node) {
	C.lxb_dom_node_insert_child(n.Ptr(), node.Ptr())
}

func (n *Node) FirstChild() *Node {
	nodePtr := C.lxb_dom_node_first_child(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{
		ptr:      nodePtr,
		document: n.document,
	}
}

func (n *Node) LastChild() *Node {
	nodePtr := C.lxb_dom_node_last_child(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{
		ptr:      nodePtr,
		document: n.document,
	}
}

func (n *Node) Parent() *Node {
	nodePtr := C.lxb_dom_node_parent(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{
		ptr:      nodePtr,
		document: n.document,
	}
}

func (n *Node) Next() *Node {
	nodePtr := C.lxb_dom_node_next(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{
		ptr:      nodePtr,
		document: n.document,
	}
}

func (n *Node) Prev() *Node {
	nodePtr := C.lxb_dom_node_prev(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{
		ptr:      nodePtr,
		document: n.document,
	}
}

func (n *Node) Children() []*Node {
	nodes := make([]*Node, 0)
	firstChild := n.FirstChild()

	if firstChild == nil {
		return nodes
	}

	nodes = append(nodes, firstChild)
	child := firstChild.Next()

	for ; child != nil; child = child.Next() {
		nodes = append(nodes, child)
	}

	return nodes
}

func (n *Node) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(n.Ptr()))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (n *Node) Element() *Element {
	return &Element{
		ptr:      (*C.lxb_dom_element_t)(unsafe.Pointer(n.ptr)),
		document: n.document,
	}
}

func (n *Node) Ptr() *C.lxb_dom_node_t {
	return n.ptr
}
