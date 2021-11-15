package html

// #include <lexbor/html/html.h>
import "C"

type Node struct {
	lexborNode *C.lxb_dom_node_t
}

func (n *Node) InsertChild(node *Node) {
	C.lxb_dom_node_insert_child(n.Ptr(), node.Ptr())
}

func (n *Node) Ptr() *C.lxb_dom_node_t {
	return n.lexborNode
}
