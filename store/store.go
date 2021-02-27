package store

// AlermStore :
type AlermStore interface {
	FindID(string, interface{}) error
	UpdateID(string, UpdateAlermParam) error
}

// BedStore :
type BedStore interface {
	FindID(string, interface{}) error
	UpdateID(string, UpdateBedParam) error
}

// UserStore :
type UserStore interface {
	FindID(string, interface{}) error
	UpdateID(string, UpdateUserParam) error
}

// UpdateAlermParam :
type UpdateAlermParam struct {
	Stop *bool
}

// UpdateBedParam :
type UpdateBedParam struct {
	Empty *bool
}

// UpdateUserParam :
type UpdateUserParam struct {
	Eye *string
}
