package html

// #include <lexbor/html/html.h>
import "C"

type HTMLElement struct {
	lexborHTMLEl *C.lxb_html_element_t
}

func (he *HTMLElement) Ptr() *C.lxb_html_element_t {
	return he.lexborHTMLEl
}
