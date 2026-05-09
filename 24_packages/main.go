package main

import (
	"fmt"

	"github.com/bedantchh/learn-go/auth"
	"github.com/bedantchh/learn-go/user"
)

func main() {
	auth.LoginWithCred("Jhon", "password")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.U{
		Email:    "user@go.com",
		Password: "user123",
	}

	fmt.Println(user.Email)
}
