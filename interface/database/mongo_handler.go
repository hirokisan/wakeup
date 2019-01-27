package database

type MongoHandler interface {
	Col(string) MongoCollectionHandler
}
