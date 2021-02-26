package query

import (
	"github.com/hirokisan/go-sample-clean-architecture/domain"
	"github.com/hirokisan/go-sample-clean-architecture/store"
)

// BedQuery :
type BedQuery struct {
	bedStore store.BedStore
}

// NewBedQuery :
func NewBedQuery(store store.BedStore) *BedQuery {
	return &BedQuery{store}
}

// BedByUserID :
func (q *BedQuery) BedByUserID(id string) (domain.Bed, error) {
	var result domain.Bed
	if err := q.bedStore.FindID(id, &result); err != nil {
		return result, err
	}
	return result, nil
}
