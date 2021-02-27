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

// Find :
func (s *AlermStore) Find(query store.Query, result interface{}) error {
	cond := bson.M{}
	// TODO : 汎用的に他のoperatorにも対応できるようにすると良さそう
	if query.Operator == store.RelationalOperatorEqual {
		cond[query.Field] = query.Value
	}
	return s.collection.Find(cond).One(&result)
}

// UpdateID :
func (s *AlermStore) UpdateID(id string, param store.UpdateAlermParam) error {
	update := bson.M{}
	if param.Stop != nil {
		update["stop"] = *param.Stop
	}
	return s.collection.UpdateId(id, bson.M{"$set": update})
}
