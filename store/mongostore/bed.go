package mongostore

import "gopkg.in/mgo.v2"

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
