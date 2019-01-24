package infra

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/hirokisan/go-sample-clean-architecture/interface/database"
)

type MongoHandler struct {
	Session    *mgo.Session
	Db         string
	Collection string
}

func NewMongoHandler(url string, db string, col string) database.MongoHandler {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	mongoHandler := new(MongoHandler)
	mongoHandler.Session = session
	mongoHandler.Db = db
	mongoHandler.Collection = col
	return mongoHandler
}

func (m *MongoHandler) Insert(dest interface{}) error {
	err := m.Session.DB(m.Db).C(m.Collection).Insert(dest)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoHandler) FindAll(res interface{}) error {
	err := m.Session.DB(m.Db).C(m.Collection).Find(nil).All(res)
	if err != nil {
		return err
	}
	return nil
}
