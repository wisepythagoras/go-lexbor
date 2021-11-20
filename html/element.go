package html

// #include <lexbor/html/html.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Element struct {
	ptr      *C.lxb_dom_element_t
	document *Document
}

func (e *Element) Attribute(attr string) string {
	cAttr := GoStringToCUChar(attr)
	attrLen := CLen(attr)
	var valLen *C.ulong

	attrVal := C.lxb_dom_element_get_attribute(e.ptr, cAttr, attrLen, valLen)

	return CUCharToGoString(attrVal)
}

func (e *Element) SetAttribute(attr string, val string) bool {
	cAttr := GoStringToCUChar(attr)
	cVal := GoStringToCUChar(val)
	attrLen := CLen(attr)
	valLen := CLen(val)

	if C.lxb_dom_element_set_attribute(e.ptr, cAttr, attrLen, cVal, valLen) == nil {
		return false
	}

	return true
}

func (e *Element) HasAttribute(attr string) bool {
	cAttr := GoStringToCUChar(attr)
	attrLen := CLen(attr)

	return (bool)(C.lxb_dom_element_has_attribute(e.ptr, cAttr, attrLen))
}

func (e *Element) FirstAttribute() *DomAttr {
	domAttr := C.lxb_dom_element_first_attribute(e.ptr)

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
	domAttr := C.lxb_dom_element_attr_by_name(e.ptr, cAttr, attrLen)

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

func (e *Element) ElementsByAttr(attr string, val string) ([]*Element, error) {
	elements := make([]*Element, 0)

	if e.document == nil {
		return elements, errors.New("No document in context")
	}

	cAttr := (*C.uchar)(unsafe.Pointer(C.CString(attr)))
	attrLen := (C.ulong)(len(attr))
	cVal := (*C.uchar)(unsafe.Pointer(C.CString(val)))
	valLen := (C.ulong)(len(val))

	// The size here should be dynamic.
	collection := CreateDomCollection(e.document, 128)
	status := C.lxb_dom_elements_by_attr(
		e.ptr,
		collection.ptr,
		cAttr,
		attrLen,
		cVal,
		valLen,
		true,
	)

	if status != C.LXB_STATUS_OK {
		return elements, errors.New("Unable to get elments by attribute")
	}

	return collection.Elements(), nil
}

func (e *Element) ElementsByTagName(tagName string) ([]*Element, error) {
	cTagName := (*C.uchar)(unsafe.Pointer(C.CString(tagName)))
	tagNameLen := (C.ulong)(len(tagName))

	// The size here should be dynamic.
	collection := CreateDomCollection(e.document, 128)
	status := C.lxb_dom_elements_by_tag_name(e.ptr, collection.ptr, cTagName, tagNameLen)

	if status != C.LXB_STATUS_OK {
		return make([]*Element, 0), errors.New("Unable to get elments by tag name")
	}

	return collection.Elements(), nil
}

func (e *Element) ElementsByClassName(className string) ([]*Element, error) {
	cClassName := (*C.uchar)(unsafe.Pointer(C.CString(className)))
	classNameLen := (C.ulong)(len(className))

	// The size here should be dynamic.
	collection := CreateDomCollection(e.document, 128)
	status := C.lxb_dom_elements_by_class_name(e.ptr, collection.ptr, cClassName, classNameLen)

	if status != C.LXB_STATUS_OK {
		return make([]*Element, 0), errors.New("Unable to get elments by tag name")
	}

	return collection.Elements(), nil
}

func (e *Element) HTMLElement() *HTMLElement {
	lxbHTMLEl := (*C.lxb_html_element_t)(unsafe.Pointer(e.ptr))
	return &HTMLElement{lexborHTMLEl: lxbHTMLEl}
}

func (e *Element) Node() *Node {
	lxbNode := (*C.lxb_dom_node_t)(unsafe.Pointer(e.ptr))
	return &Node{lexborNode: lxbNode}
}
