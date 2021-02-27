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

// Find :
func (s *UserStore) Find(query store.Query, result interface{}) error {
	cond := bson.M{}
	// TODO : 汎用的に他のoperatorにも対応できるようにすると良さそう
	if query.Operator == store.RelationalOperatorEqual {
		cond[query.Field] = query.Value
	}
	return s.collection.Find(cond).One(&result)
}

// UpdateID :
func (s *UserStore) UpdateID(id string, param store.UpdateUserParam) error {
	update := bson.M{}
	if param.Eye != nil {
		update["eye"] = *param.Eye
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
