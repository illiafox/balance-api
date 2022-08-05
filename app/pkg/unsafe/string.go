package unsafe

import (
	"reflect"
	"unsafe"
)

//go:nosplit
func BytesToString(b []byte) (s string) {
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))

	str.Data = slice.Data
	str.Len = slice.Len

	return
}
