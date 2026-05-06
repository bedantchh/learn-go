package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{1, 2, 3}
	// var sum int
	// for _, val := range nums {
	// 	sum = sum + val
	// }
	// fmt.Println(sum)

	// m := map[string]string{"name": "ubuntu", "version": "22.4"}

	// for key, val := range m {
	// 	// fmt.Println(key, val)
	// 	fmt.Printf("%v : %v \n", key, val)
	// }

	// for val := range m {
	// 	fmt.Println(val)
	// 	// fmt.Printf("%v : %v \n", key, val)
	// }

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
