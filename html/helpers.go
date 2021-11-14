package html

// #include <stdio.h>
import "C"
import (
	"unsafe"
)

func GoStringToCUChar(str string) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(C.CString(str)))
}

func CUCharToGoString(str *C.uchar) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

func CLen(str string) C.ulong {
	return (C.ulong)(len(str))
}
