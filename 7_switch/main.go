package main

import "fmt"

func main() {
	// simple switch, no need break
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("Other")
	// }

	//multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	{
	// 		fmt.Println("workday")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Workday")
	// 	}
	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("str")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI(8)
}
