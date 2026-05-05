package main

import "fmt"

func main(){
	//while loop in go
	// var i = 0
	// for i<=10{
	// 	println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	println("hehe")
	// }


	//for loop (continue abd break is supported too)
	// for i:=0;i<=3;i++ {
	// 	fmt.Println(i)
	// }

	//range
	for i := range 3 {
		fmt.Println(i)
	}
}
