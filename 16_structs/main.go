package main

import (
	"fmt"
	"time"
)

// order struct
type cutomer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	cutomer
}

// func newOrder(id string, amount float32, status string) *order {
// 	//initial setup goes here
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// receiver type
// func (o *order) statusChange(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	// myCustomer := cutomer{
	// 	name:  "Jhon",
	// 	phone: "98989",
	// }
	newOrder := order{
		id:     "1",
		amount: 50.50,
		status: "Received",
		cutomer: cutomer{
			name:  "Jhon",
			phone: "98989",
		},
	}
	fmt.Println(newOrder)

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.50,
	// 	status: "Received",
	// }

	// myOrder.statusChange("Ordered")
	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder)
	// amt := myOrder.getAmount()
	// fmt.Println(amt)

	// myNewOrder := newOrder("1", 50.25, "Ordered")
	// fmt.Println(myNewOrder.amount)
}
