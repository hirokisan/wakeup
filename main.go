package main

import (
	"github.com/hirokisan/go-sample-clean-architecture/database/query"
	"github.com/hirokisan/go-sample-clean-architecture/store/mongostore"
	"github.com/hirokisan/go-sample-clean-architecture/usecase"
)

func main() {
	store, teardown := mongostore.New("", "")
	defer teardown()

	alermQuery := query.NewAlermQuery(store.AlermStore())
	bedQuery := query.NewBedQuery(store.BedStore())

	u := usecase.NewWakeUpUsecase(
		alermQuery,
		bedQuery,
	)
	if err := u.WakeUp; err != nil {
		panic(err)
	}
}
