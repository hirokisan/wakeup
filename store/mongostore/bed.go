package mongostore

import (
	"github.com/hirokisan/go-sample-clean-architecture/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BedStore :
type BedStore struct {
	collection *mgo.Collection
}

// BedStore :
func (s *MongoStore) BedStore() *BedStore {
	return &BedStore{s.db.C("beds")}
}

// Find :
func (s *BedStore) Find(query store.Query, result interface{}) error {
	cond := bson.M{}
	// TODO : 汎用的に他のoperatorにも対応できるようにすると良さそう
	if query.Operator == store.RelationalOperatorEqual {
		cond[query.Field] = query.Value
	}
	return s.collection.Find(cond).One(&result)
}

// UpdateID :
func (s *BedStore) UpdateID(id string, param store.UpdateBedParam) error {
	update := bson.M{}
	if param.Empty != nil {
		update["empty"] = *param.Empty
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
