package html

/*
#include <lexbor/html/html.h>
#include <lexbor/html/parser.h>
#include <lexbor/html/serialize.h>
#include <lexbor/html/interfaces/document.h>
#include <lexbor/dom/interface.h>

lxb_dom_node_t *lxb_dom_interface_node_custom(lxb_html_document_t *obj) {
	return (lxb_dom_node_t *) (obj);
}

lxb_inline lxb_status_t
serializer_callback(const lxb_char_t *data, size_t len, void *ctx)
{
    printf("%.*s", (int) len, (const char *) data);

    return LXB_STATUS_OK;
}

lxb_inline void
serialize(lxb_dom_node_t *node)
{
    lxb_status_t status;

    status = lxb_html_serialize_pretty_tree_cb(node,
                                               LXB_HTML_SERIALIZE_OPT_UNDEF,
                                               0, serializer_callback, NULL);
    if (status != LXB_STATUS_OK) {
        printf("Failed to serialization HTML tree\n");
    }
}
*/
import "C"
import (
	"unsafe"
)

type Document struct {
	lexborDoc *C.lxb_html_document_t
}

func (d *Document) Create() {
	document := C.lxb_html_document_create()
	d.lexborDoc = document
}

func (d *Document) Parse(html string) bool {
	if d.lexborDoc == nil {
		return false
	}

	unsafeHtml := unsafe.Pointer(C.CString(html))
	status := C.lxb_html_document_parse(d.lexborDoc, (*C.uchar)(unsafeHtml), (C.ulong)(len(html)))

	if status != C.LXB_STATUS_OK {
		return false
	}

	return true
}

func (d *Document) BodyElement() *BodyElement {
	bodyElement := C.lxb_html_document_body_element(d.lexborDoc)
	body := &BodyElement{
		lexborElement: bodyElement,
	}

	return body
}

func (d *Document) DomInterfaceNode() *Node {
	lexborNode := (*C.lxb_dom_node_t)(unsafe.Pointer(d.lexborDoc))

	// An alternative way:
	// lexborNode := C.lxb_dom_interface_node_custom(d.lexborDoc)

	node := &Node{
		lexborNode: lexborNode,
	}

	return node
}

func (d *Document) Destroy() {
	C.lxb_html_document_destroy(d.lexborDoc)
}

func Serialize(node *Node) {
	C.serialize(node.Ptr())
}
