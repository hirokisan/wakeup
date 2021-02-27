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

// FindID :
func (s *BedStore) FindID(id string, result interface{}) error {
	return s.collection.FindId(id).One(&result)
}

// UpdateID :
func (s *BedStore) UpdateID(id string, param store.UpdateBedParam) error {
	update := bson.M{}
	if param.Empty != nil {
		update["empty"] = *param.Empty
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
