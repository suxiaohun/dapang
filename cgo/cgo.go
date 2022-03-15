package cgo

/*
#cgo CFLAGS: -I .
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "hello.h"
void hey(_GoString_ s);


*/
import "C"

import (
	"fmt"
	"unsafe"
)

func HelloWorldWithCMethod() {
	cs1 := C.CString("cgo: Hello World")
	C.free(unsafe.Pointer(cs1))
	C.puts(cs1)
}

//export say
func say(s *C.char) {
	fmt.Println(C.GoString(s))
}

func Say() {
	C.say(C.CString("cgo: say"))
}

//export hey
func hey(s string)  {
	fmt.Println(s)
}

func Hey()  {
	C.hey("cgo: hey")
}
