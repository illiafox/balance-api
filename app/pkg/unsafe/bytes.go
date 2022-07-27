package unsafe

import (
	"reflect"
	"unsafe"
)

func StringToBytes(s *string) []byte {
	ptr := unsafe.Pointer(s)
	slice := *(*[]byte)(ptr)
	l := (*reflect.StringHeader)(ptr).Len
	return slice[:l:l]
}
