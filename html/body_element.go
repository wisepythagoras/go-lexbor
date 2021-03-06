package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type BodyElement struct {
	ptr      *C.lxb_html_body_element_t
	document *Document
}

func (b *BodyElement) Element() *Element {
	element := &Element{
		ptr:      (*C.lxb_dom_element_t)(unsafe.Pointer(b.ptr)),
		document: b.document,
	}

	return element
}

func (b *BodyElement) Ptr() *C.lxb_html_body_element_t {
	return b.ptr
}
