package query

import "github.com/hirokisan/go-sample-clean-architecture/domain"

// BedQuery :
type BedQuery struct {
	// store
}

// NewBedQuery :
func NewBedQuery() *BedQuery {
	return &BedQuery{}
}

// BedByUserID :
func (q *BedQuery) BedByUserID(id string) (domain.Bed, error) {
	var result domain.Bed
	// todo : implement
	return result, nil
}
