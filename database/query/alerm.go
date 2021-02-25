package query

import "github.com/hirokisan/go-sample-clean-architecture/domain"

// AlermQuery :
type AlermQuery struct {
	// store
}

// NewAlermQuery :
func NewAlermQuery() *AlermQuery {
	return &AlermQuery{}
}

// AlermByUserID :
func (q *AlermQuery) AlermByUserID(id string) (domain.Alerm, error) {
	var result domain.Alerm
	// todo : implement
	return result, nil
}
