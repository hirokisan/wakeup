package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

type CollectionHandler struct {
	Collection *mgo.Collection
}

func (h *CollectionHandler) Insert(dest interface{}) error {
	err := h.Collection.Insert(dest)
	if err != nil {
		return err
	}
	return nil
}

func (h *CollectionHandler) FindAll(res interface{}) error {
	err := h.Collection.Find(nil).All(res)
	if err != nil {
		return err
	}
	return nil
}
