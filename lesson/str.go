package lesson

import "fmt"

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	fmt.Printf("string reverse: %v=>%v\n", s, string(b))
	return string(b)
}
