package html

// #include <lexbor/html/html.h>
import "C"

type BodyElement struct {
	lexborElement *C.lxb_html_body_element_t
}

func (b *BodyElement) Ptr() *C.lxb_html_body_element_t {
	return b.lexborElement
}
