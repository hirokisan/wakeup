package mongostore

import (
	"github.com/hirokisan/go-sample-clean-architecture/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AlermStore :
type AlermStore struct {
	collection *mgo.Collection
}

// AlermStore :
func (s *MongoStore) AlermStore() *AlermStore {
	return &AlermStore{s.db.C("alerms")}
}

// FindID :
func (s *AlermStore) FindID(id string, result interface{}) error {
	return s.collection.FindId(id).One(&result)
}

// UpdateID :
func (s *AlermStore) UpdateID(id string, param store.UpdateAlermParam) error {
	update := bson.M{}
	if param.Stop != nil {
		update["stop"] = *param.Stop
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
