package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type Collection struct {
	lxbCollection *C.lxb_dom_collection_t
}

func (c *Collection) Length() int {
	return int(c.lxbCollection.array.length)
}

func (c *Collection) Size() int {
	return int(c.lxbCollection.array.size)
}

func (c *Collection) Element(idx int) *Element {
	if idx >= c.Length() {
		return nil
	}

	return &Element{
		lexborElement: C.lxb_dom_collection_element(c.lxbCollection, (C.ulong)(idx)),
	}
}

func (c *Collection) DomElementsByTagName(tagName string, el *Element) []*Element {
	cTagName := (*C.uchar)(unsafe.Pointer(C.CString(tagName)))
	tagNameLen := (C.ulong)(len(tagName))
	elements := make([]*Element, 0)
	status := C.lxb_dom_elements_by_tag_name(el.Ptr(), c.lxbCollection, cTagName, tagNameLen)

	if status != C.LXB_STATUS_OK || c.Length() == 0 {
		return elements
	}

	for i := 0; i < c.Length(); i++ {
		elements = append(elements, c.Element(i))
	}

	return elements
}

func (c *Collection) Destroy() {
	C.lxb_dom_collection_destroy(c.lxbCollection, true)
}

func CreateDomCollection(doc *Document, size int) *Collection {
	lxbCollection := C.lxb_dom_collection_make(doc.DomDocument(), (C.ulong)(size))

	if lxbCollection == nil {
		return nil
	}

	return &Collection{lxbCollection: lxbCollection}
}
