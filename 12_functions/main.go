package main

// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string) {
// 	return "golang", "javascript", "C"
// }

// func processInt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 3
	}
}

func main() {
	// x := add(2, 3)
	// fmt.Println(x)
	// fmt.Println(getLanguages())

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	// fn := func(a int) int {
	// 	return 2
	// }
	fn := processIt()
	fn(6)
}
