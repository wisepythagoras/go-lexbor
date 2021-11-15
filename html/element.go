package html

// #include <lexbor/html/html.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Element struct {
	lexborElement *C.lxb_dom_element_t
}

func (e *Element) Attribute(attr string) string {
	cAttr := GoStringToCUChar(attr)
	attrLen := CLen(attr)
	var valLen *C.ulong

	attrVal := C.lxb_dom_element_get_attribute(e.Ptr(), cAttr, attrLen, valLen)

	return CUCharToGoString(attrVal)
}

func (e *Element) SetAttribute(attr string, val string) bool {
	cAttr := GoStringToCUChar(attr)
	cVal := GoStringToCUChar(val)
	attrLen := CLen(attr)
	valLen := CLen(val)

	if C.lxb_dom_element_set_attribute(e.Ptr(), cAttr, attrLen, cVal, valLen) == nil {
		return false
	}

	return true
}

func (e *Element) HasAttribute(attr string) bool {
	cAttr := GoStringToCUChar(attr)
	attrLen := CLen(attr)

	return (bool)(C.lxb_dom_element_has_attribute(e.Ptr(), cAttr, attrLen))
}

func (e *Element) FirstAttribute() *DomAttr {
	domAttr := C.lxb_dom_element_first_attribute(e.Ptr())

	if domAttr == nil {
		return nil
	}

	return &DomAttr{lexborDomAttr: domAttr}
}

func (e *Element) NextAttribute(attr *DomAttr) *DomAttr {
	domAttr := C.lxb_dom_element_next_attribute(attr.Ptr())

	if domAttr == nil {
		return nil
	}

	return &DomAttr{lexborDomAttr: domAttr}
}

func (e *Element) AttributeByName(attr string) *DomAttr {
	cAttr := GoStringToCUChar(attr)
	attrLen := CLen(attr)
	domAttr := C.lxb_dom_element_attr_by_name(e.Ptr(), cAttr, attrLen)

	if domAttr == nil {
		return nil
	}

	return &DomAttr{lexborDomAttr: domAttr}
}

func (e *Element) SetInnerHTML(innerHTML string) error {
	cInner := GoStringToCUChar(innerHTML)
	innerLen := CLen(innerHTML)
	el := C.lxb_html_element_inner_html_set(e.HTMLElement().Ptr(), cInner, innerLen)

	if el == nil {
		return errors.New("Failed to parse inner HTML")
	}

	return nil
}

func (e *Element) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(e.Ptr()))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (e *Element) Node() *Node {
	lxbNode := (*C.lxb_dom_node_t)(unsafe.Pointer(e.Ptr()))
	return &Node{lexborNode: lxbNode}
}

func (e *Element) Ptr() *C.lxb_dom_element_t {
	return e.lexborElement
}
