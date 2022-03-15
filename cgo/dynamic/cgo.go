package dynamic

//使用#cgo定义库路径

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ./libs -ldynamic_hello
#include "hello.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Hello2() {
	C.StaticHello()
	C.DynamicHello()
}



func Hello3()  {

	//str := "aaa"
	//var a = C.CString(str)





	//fmt.Printf("=====%T===\n",str)
	//fmt.Printf("=====%T===\n",a)

	name := C.Name()

	//fmt.Printf("=====%T===\n",name)
	//fmt.Printf("=====%s===\n",&name)

	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(C.GoString(name))


	C.DynamicHello()
}

func charpp2string(charpp **C.char, n int) []string {

	var b *C.char
	ptrSize := unsafe.Sizeof(b)
	gostring := make([]string, n)
	if n > 0 {
		for i := 0; i < n; i++ {

			element := (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(charpp)) + uintptr(i)*ptrSize))
			gostring[i] = C.GoString((*C.char)(*element))
		}

		//(*C.char)(*(**C.char)(unsafe.Pointer(uintptr( unsafe.Pointer(job.exHosts)) + uintptr(1)*ptrSize ) ) ))
	}

	return gostring
}