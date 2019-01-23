package main

import (
	"fmt"

	"github.com/hirokisan/clean/domain"
	"github.com/hirokisan/clean/infra"
)

func main() {
	handler := infra.NewMongoHandler("local", "user")
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
