package html

// #include <lexbor/html/html.h>
import "C"

type TagHash struct {
	lexborTagHash *C.lexbor_hash_t
}

func (th *TagHash) Ptr() *C.lexbor_hash_t {
	return th.lexborTagHash
}
