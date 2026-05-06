package main

import "fmt"

// by value
// func changeNum(num int) int {
// 	num = 999
// 	return num
// }

// by refrence
func changeNum(num *int) int {
	*num = 999
	return *num
}

func main() {
	num := 666
	hehe := changeNum(&num)
	fmt.Println(num)
	fmt.Println(hehe)
}
