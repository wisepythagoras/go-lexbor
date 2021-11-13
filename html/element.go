package html

// #include <lexbor/html/html.h>
import "C"
import (
	"unsafe"
)

type Element struct {
	lexborElement *C.lxb_dom_element_t
}

func (e *Element) Attribute(attr string) string {
	cAttr := (*C.uchar)(unsafe.Pointer(C.CString(attr)))
	attrLen := (C.ulong)(len(attr))
	var valLen *C.ulong

	attrVal := C.lxb_dom_element_get_attribute(e.Ptr(), cAttr, attrLen, valLen)

	return C.GoString((*C.char)(unsafe.Pointer(attrVal)))
}

func (e *Element) Ptr() *C.lxb_dom_element_t {
	return e.lexborElement
}
