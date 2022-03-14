package lesson

import (
	"encoding/json"
	"fmt"
)

func Map2JSON() {
	var s []map[string]interface{}

	m1 := map[string]interface{}{"name": "John", "age": 10}
	m2 := map[string]interface{}{"name": "Alex", "age": 12}

	s = append(s, m1, m2)
	s = append(s, m2)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println("Map2JSON:", string(b))

}
