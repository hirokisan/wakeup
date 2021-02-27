package main

import (
	"github.com/hirokisan/go-sample-clean-architecture/database/command"
	"github.com/hirokisan/go-sample-clean-architecture/database/query"
	"github.com/hirokisan/go-sample-clean-architecture/store/mongostore"
	"github.com/hirokisan/go-sample-clean-architecture/usecase"
)

func main() {
	store, teardown := mongostore.New("", "")
	defer teardown()

	userCommand := command.NewUserCommand(store.UserStore())
	bedCommand := command.NewBedCommand(store.BedStore())
	alermCommand := command.NewAlermCommand(store.AlermStore())
	alermQuery := query.NewAlermQuery(store.AlermStore())
	bedQuery := query.NewBedQuery(store.BedStore())

	u := usecase.NewWakeUpUsecase(
		*userCommand,
		*bedCommand,
		*alermCommand,
		*alermQuery,
		*bedQuery,
	)
	if err := u.WakeUp; err != nil {
		panic(err)
	}
}
