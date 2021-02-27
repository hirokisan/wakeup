package store

// AlermStore :
type AlermStore interface {
	Find(Query, interface{}) error
	UpdateID(string, UpdateAlermParam) error
}

// BedStore :
type BedStore interface {
	Find(Query, interface{}) error
	UpdateID(string, UpdateBedParam) error
}

// UserStore :
type UserStore interface {
	Find(Query, interface{}) error
	UpdateID(string, UpdateUserParam) error
}
