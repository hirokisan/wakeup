package mongostore

import "gopkg.in/mgo.v2"

// MongoStore :
type MongoStore struct {
	db *mgo.Database
}

// New :
func New(url string, dbname string) (*MongoStore, func()) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return &MongoStore{session.DB(dbname)}, session.Close
}
