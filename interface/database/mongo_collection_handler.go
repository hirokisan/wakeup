package database

type MongoCollectionHandler interface {
	Insert(interface{}) error
	FindAll(interface{}) error
}
