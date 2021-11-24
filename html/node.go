package html

// #include <lexbor/html/html.h>
import "C"
import "unsafe"

type Node struct {
	ptr *C.lxb_dom_node_t
}

func (n *Node) InsertChild(node *Node) {
	C.lxb_dom_node_insert_child(n.Ptr(), node.Ptr())
}

func (n *Node) FirstChild() *Node {
	nodePtr := C.lxb_dom_node_first_child(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{ptr: nodePtr}
}

func (n *Node) LastChild() *Node {
	nodePtr := C.lxb_dom_node_last_child(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{ptr: nodePtr}
}

func (n *Node) Parent() *Node {
	nodePtr := C.lxb_dom_node_parent(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{ptr: nodePtr}
}

func (n *Node) Next() *Node {
	nodePtr := C.lxb_dom_node_next(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{ptr: nodePtr}
}

func (n *Node) Prev() *Node {
	nodePtr := C.lxb_dom_node_prev(n.Ptr())

	if nodePtr == nil {
		return nil
	}

	return &Node{ptr: nodePtr}
}

func (n *Node) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(n.Ptr()))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (n *Node) Ptr() *C.lxb_dom_node_t {
	return n.ptr
}
