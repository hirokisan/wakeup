package query

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
	"github.com/hirokisan/go-sample-clean-architecture/store"
)

// AlermQuery :
type AlermQuery struct {
	alermStore store.AlermStore
}

// NewAlermQuery :
func NewAlermQuery(store store.AlermStore) *AlermQuery {
	return &AlermQuery{store}
}

// AlermByUserID :
func (q *AlermQuery) AlermByUserID(userID string) (domain.Alerm, error) {
	var result domain.Alerm
	if err := q.alermStore.Find(store.Query{
		Field:    "userId",
		Operator: store.RelationalOperatorEqual,
		Value:    userID,
	}, &result); err != nil {
		return result, err
	}
	return result, nil
}
