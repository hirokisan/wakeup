package database

type MongoHandler interface {
	Insert(interface{}) error
	FindAll(interface{}) error
}
