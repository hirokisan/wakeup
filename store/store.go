package store

// AlermStore :
type AlermStore interface {
	FindID(string, interface{}) error
	UpdateID(string, UpdateAlermParam) error
}

// BedStore :
type BedStore interface {
	FindID(string, interface{}) error
}

// UpdateAlermParam :
type UpdateAlermParam struct {
	Stop *bool
}
