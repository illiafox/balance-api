package unsafe

import (
	"reflect"
	"unsafe"
)

//go:nosplit
func StringToBytes(s *string) (b []byte) {
	str := (*reflect.StringHeader)(unsafe.Pointer(s))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	slice.Data = str.Data
	slice.Len = str.Len
	slice.Cap = str.Len

	return
}
