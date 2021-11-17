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
	"errors"
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

func (d *Document) ParseChunks(htmlChunks []string) bool {
	status := C.lxb_html_document_parse_chunk_begin(d.lexborDoc)

	if status != C.LXB_STATUS_OK {
		return false
	}

	for _, chunk := range htmlChunks {
		uCharChunk := (*C.uchar)(unsafe.Pointer(C.CString(chunk)))
		status = C.lxb_html_document_parse_chunk(d.lexborDoc, uCharChunk, (C.ulong)(len(chunk)))

		if status != C.LXB_STATUS_OK {
			return false
		}
	}

	return true
}

func (d *Document) BodyElement() *BodyElement {
	bodyElement := C.lxb_html_document_body_element(d.lexborDoc)
	body := &BodyElement{
		lexborElement: bodyElement,
		document:      d,
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

func (d *Document) Title() string {
	if d.lexborDoc == nil {
		return ""
	}

	var title_len C.size_t
	title := C.lxb_html_document_title(d.lexborDoc, &title_len)

	return C.GoString((*C.char)(unsafe.Pointer(title)))
}

func (d *Document) ChangeTitle(title string) bool {
	uCharPtrTitle := (*C.uchar)(unsafe.Pointer(C.CString(title)))
	titleLen := (C.ulong)(len(title))
	status := C.lxb_html_document_title_set(d.lexborDoc, uCharPtrTitle, titleLen)

	if status != C.LXB_STATUS_OK {
		return false
	}

	return true
}

func (d *Document) Tags() (*TagHash, []string, map[string]bool) {
	// This will contain all tag names.
	tagNames := make([]string, 0)

	// The tag state will be either true or false based on whether it's a void element
	tagStates := make(map[string]bool)

	lxbTags := C.lxb_html_document_tags(d.Ptr())

	for tagId := (int)(C.LXB_TAG_A); tagId < (int)(C.LXB_TAG__LAST_ENTRY); tagId++ {
		var tagNameLen *C.ulong
		tagName := C.lxb_tag_name_by_id(lxbTags, (C.ulong)(tagId), tagNameLen)
		tagNames = append(tagNames, CUCharToGoString(tagName))
		tagStates[CUCharToGoString(tagName)] = (bool)(C.lxb_html_tag_is_void((C.ulong)(tagId)))
	}

	return &TagHash{lexborTagHash: lxbTags}, tagNames, tagStates
}

func (d *Document) CreateElement(name string) *Element {
	cName := GoStringToCUChar(name)
	nameLen := CLen(name)

	lxbElement := C.lxb_dom_document_create_element(d.DomDocument(), cName, nameLen, nil)

	if lxbElement == nil {
		return nil
	}

	return &Element{
		lexborElement: lxbElement,
		document:      d,
	}
}

func (d *Document) CreateTextNode(text string) *DomText {
	cText := GoStringToCUChar(text)
	textLen := CLen(text)

	domText := C.lxb_dom_document_create_text_node(d.DomDocument(), cText, textLen)

	if domText == nil {
		return nil
	}

	return &DomText{lexborDomText: domText}
}

func (d *Document) GetElementById(id string) (*Element, error) {
	elements, err := d.BodyElement().Element().ElementsByAttr("id", id)

	if err != nil {
		return nil, err
	}

	if len(elements) == 0 {
		return nil, errors.New("Element with this id not found")
	}

	return elements[0], nil
}

func (d *Document) GetElementsByTagName(tagName string) ([]*Element, error) {
	return d.BodyElement().Element().ElementsByTagName(tagName)
}

func (d *Document) DomDocument() *C.lxb_dom_document_t {
	if d.lexborDoc == nil {
		return nil
	}

	return &d.lexborDoc.dom_document
}

func (d *Document) Destroy() {
	C.lxb_html_document_destroy(d.lexborDoc)
}

func (d *Document) Ptr() *C.lxb_html_document_t {
	return d.lexborDoc
}

func Serialize(node *Node) {
	C.serialize(node.Ptr())
}
