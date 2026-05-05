package main

import "fmt"

const age = 30 //allowed
//age := 30  //shorthand not allowed in global

func main(){
	const name = "Joe"

	// println(age)


	const (
		port = 4000
		host = "localhost"
	)

	fmt.Println(port, host)
}