package lesson

import (
	"fmt"
)

func ForStandard(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("ForStandard: ", i+1)
	}
}

func ForWithCondition(n int) {
	for n < 5 {
		n++
		fmt.Printf("ForWithCondition: %d\n", n)
	}
}

func ForInfinite() {
	n := 0
	for {
		n++
		if n > 5 {
			break
		}
		fmt.Println("ForInfinite: ", n)
	}
}

func ForRange() {
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Printf("ForRange: %d\n", v)
	}
}
