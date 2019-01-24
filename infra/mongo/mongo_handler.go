package mongo

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/hirokisan/go-sample-clean-architecture/interface/database"
)

type Handler struct {
	Session    *mgo.Session
	Db         string
	Collection string
}

func NewHandler(url string, db string, col string) database.MongoHandler {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Handler := new(Handler)
	Handler.Session = session
	Handler.Db = db
	Handler.Collection = col
	return Handler
}

func (m *Handler) Insert(dest interface{}) error {
	err := m.Session.DB(m.Db).C(m.Collection).Insert(dest)
	if err != nil {
		return err
	}
	return nil
}

func (m *Handler) FindAll(res interface{}) error {
	err := m.Session.DB(m.Db).C(m.Collection).Find(nil).All(res)
	if err != nil {
		return err
	}
	return nil
}
