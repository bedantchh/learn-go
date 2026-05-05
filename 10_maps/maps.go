package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict
func main() {
	// m := make(map[int]string)

	// m[1] = "golang"
	// m[2] = "js"

	// fmt.Println(m[1], m[2])
	// fmt.Println(len(m))

	// delete(m, 2)
	// fmt.Println(m)

	// m := map[string]int{"price": 50, "rating": 4}
	// fmt.Println(m)

	// m := map[string]int{"price": 50, "rating": 4}

	// val, ok := m["rating"]
	// fmt.Println(val)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{"price": 50, "rating": 4}
	m2 := map[string]int{"price": 50, "rating": 4}

	fmt.Println(maps.Equal(m1, m2))
}
