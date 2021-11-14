package html

// #include <lexbor/html/html.h>
import "C"

type DomAttr struct {
	lexborDomAttr *C.lxb_dom_attr_t
}

func (d *DomAttr) QualifiedName() string {
	var nameLen *C.ulong
	return CUCharToGoString(C.lxb_dom_attr_qualified_name(d.Ptr(), nameLen))
}

func (d *DomAttr) Value() string {
	var valLen *C.ulong
	return CUCharToGoString(C.lxb_dom_attr_value(d.Ptr(), valLen))
}

func (d *DomAttr) SetValue(val string) bool {
	cVal := GoStringToCUChar(val)
	valLen := CLen(val)

	status := C.lxb_dom_attr_set_value(d.Ptr(), cVal, valLen)

	if status != C.LXB_STATUS_OK {
		return false
	}

	return true
}

func (d *DomAttr) Ptr() *C.lxb_dom_attr_t {
	return d.lexborDomAttr
}
