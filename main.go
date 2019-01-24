package main

import (
	"fmt"

	"github.com/hirokisan/go-sample-clean-architecture/domain"
	"github.com/hirokisan/go-sample-clean-architecture/infra"
)

func main() {
	handler := infra.NewMongoHandler("localhost:27017", "local", "user")
	user := new(domain.User)
	user.ID = int(1)
	user.FirstName = "kenta"
	user.LastName = "yamada"
	err := handler.Insert(user)
	if err != nil {
		panic(err)
	}

	var res domain.Users
	err = handler.FindAll(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
