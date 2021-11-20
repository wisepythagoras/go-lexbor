package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type Collection struct {
	ptr      *C.lxb_dom_collection_t
	document *Document
}

func (c *Collection) Length() int {
	return int(c.ptr.array.length)
}

func (c *Collection) Size() int {
	return int(c.ptr.array.size)
}

func (c *Collection) Element(idx int) *Element {
	if idx >= c.Length() {
		return nil
	}

	return &Element{
		ptr:      C.lxb_dom_collection_element(c.ptr, (C.ulong)(idx)),
		document: c.document,
	}
}

func (c *Collection) DomElementsByTagName(tagName string, el *Element) []*Element {
	cTagName := (*C.uchar)(unsafe.Pointer(C.CString(tagName)))
	tagNameLen := (C.ulong)(len(tagName))
	elements := make([]*Element, 0)
	status := C.lxb_dom_elements_by_tag_name(el.ptr, c.ptr, cTagName, tagNameLen)

	if status != C.LXB_STATUS_OK || c.Length() == 0 {
		return elements
	}

	for i := 0; i < c.Length(); i++ {
		elements = append(elements, c.Element(i))
	}

	return elements
}

func (c *Collection) DomElementsByAttr(attr string, val string, el *Element) []*Element {
	cAttr := (*C.uchar)(unsafe.Pointer(C.CString(attr)))
	attrLen := (C.ulong)(len(attr))
	cVal := (*C.uchar)(unsafe.Pointer(C.CString(val)))
	valLen := (C.ulong)(len(val))
	elements := make([]*Element, 0)
	status := C.lxb_dom_elements_by_attr(
		el.ptr,
		c.ptr,
		cAttr,
		attrLen,
		cVal,
		valLen,
		true,
	)

	if status != C.LXB_STATUS_OK || c.Length() == 0 {
		return elements
	}

	for i := 0; i < c.Length(); i++ {
		elements = append(elements, c.Element(i))
	}

	return elements
}

func (c *Collection) Elements() []*Element {
	elements := make([]*Element, 0)

	for i := 0; i < c.Length(); i++ {
		elements = append(elements, c.Element(i))
	}

	return elements
}

func (c *Collection) Destroy() {
	C.lxb_dom_collection_destroy(c.ptr, true)
}

func CreateDomCollection(doc *Document, size int) *Collection {
	ptr := C.lxb_dom_collection_make(doc.DomDocument(), (C.ulong)(size))

	if ptr == nil {
		return nil
	}

	return &Collection{
		ptr:      ptr,
		document: doc,
	}
}
