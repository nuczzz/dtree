package dtree

import (
	"unsafe"
)

func Bytes2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func String2Bytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// IndexRight return the first index of 'b' in bytes 'src'.
// If not found, -1 returned.
func IndexRight(src []byte, b byte) int {
	for index := len(src) - 1; index >= 0; index-- {
		if src[index] == b {
			return index
		}
	}
	return -1
}
