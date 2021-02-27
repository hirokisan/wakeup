package mongostore

import (
	"github.com/hirokisan/go-sample-clean-architecture/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserStore :
type UserStore struct {
	collection *mgo.Collection
}

// UserStore :
func (s *MongoStore) UserStore() *UserStore {
	return &UserStore{s.db.C("users")}
}

// FindID :
func (s *UserStore) FindID(id string, result interface{}) error {
	return s.collection.FindId(id).One(&result)
}

// UpdateID :
func (s *UserStore) UpdateID(id string, param store.UpdateUserParam) error {
	update := bson.M{}
	if param.Eye != nil {
		update["eye"] = *param.Eye
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
