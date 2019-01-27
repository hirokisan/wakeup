package mongo

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/hirokisan/go-sample-clean-architecture/interface/database"
)

type Handler struct {
	Session *mgo.Session
	Db      string
}

func NewHandler(url string, db string) database.MongoHandler {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Handler := new(Handler)
	Handler.Session = session
	Handler.Db = db
	return Handler
}

func (m *Handler) Col(col string) database.MongoCollectionHandler {
	return &CollectionHandler{m.Session.DB(m.Db).C(col)}
}
