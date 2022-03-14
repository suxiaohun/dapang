package cgo

// 直接写C源码来调用

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"unsafe"
)

func HelloWorldWithCMethod() {
	cs := C.CString("C printf: Hello World\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
