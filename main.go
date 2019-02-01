package main

import (
	"fmt"

	"github.com/hirokisan/go-sample-clean-architecture/domain"
	"github.com/hirokisan/go-sample-clean-architecture/infra/mongo"
	"github.com/hirokisan/go-sample-clean-architecture/interface/database"
	"github.com/hirokisan/go-sample-clean-architecture/usecase"
)

func main() {
	handler := mongo.NewHandler("localhost:27017", "local")
	userHandler := &usecase.UserInteractor{
		UserRepository: &database.UserRepository{
			MongoCollectionHandler: handler.Col("user"),
		},
	}
	user := &domain.User{
		ID:        1,
		FirstName: "kenta",
		LastName:  "yamada",
	}
	err := userHandler.Add(user)
	if err != nil {
		panic(err)
	}

	var res domain.Users
	err = userHandler.Users(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
