package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type BodyElement struct {
	lexborElement *C.lxb_html_body_element_t
}

func (b *BodyElement) Element() *Element {
	element := &Element{
		lexborElement: (*C.lxb_dom_element_t)(unsafe.Pointer(b.lexborElement)),
	}

	return element
}

func (b *BodyElement) Ptr() *C.lxb_html_body_element_t {
	return b.lexborElement
}