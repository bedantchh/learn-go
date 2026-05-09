package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())
	// fmt.Println("Folder:", fileInfo.IsDir())
	// fmt.Println("File Size:", fileInfo.Size())
	// fmt.Println("File Permission:", fileInfo.Mode())
	// fmt.Println("File modified at:", fileInfo.ModTime())

	//read file content
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 16)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// fmt.Println("data", d, buf)

	//simple way (loads in memory)
	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	//read folders
	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("hi golang")
	// bytes := []byte("hello golanggg")
	// f.Write(bytes)

	// read and write to another file - streaming
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()
	fmt.Println("written to new file")

	// os.Remove("example2.txt")
}
