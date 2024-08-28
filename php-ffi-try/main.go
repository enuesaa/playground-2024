package main

import "C"

//export countTextLength
func countTextLength(cstr *C.char) C.int {
	text := C.GoString(cstr)
	length := len(text)

	return C.int(length) 
}

func main() {}
