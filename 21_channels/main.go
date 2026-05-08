package main

import (
	"fmt"
	"time"
)

//sending
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number:", num)
// 		time.Sleep(time.Second)
// 	}

// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

//goroutine synchronizer
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing")

// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"

	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 3; i++ {
	// 	emailChan <- fmt.Sprintf("%d@go.com", i)
	// }

	// fmt.Println("done sending")
	// // emailChan <- "1@example.com"
	// // emailChan <- "2@example.com"

	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)
	// close(emailChan)
	// <-done

	// done := make(chan bool)

	// go task(done)
	// <-done

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)
	// numChan <- 5

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// msg := <-messageChan

	// fmt.Println(msg)
}
