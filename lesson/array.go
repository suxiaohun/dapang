package lesson

import "fmt"

func Array1() {
	var a [3]int
	//var a = [...]int{99: -1}
	fmt.Println("array[0]: ", a[0])
	fmt.Println("array len: ", len(a))

	for i, v := range a {
		fmt.Printf("range array: %d %v\n", i, v)
	}

}
