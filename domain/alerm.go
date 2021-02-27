package domain

// Alerm :
type Alerm struct {
	ID     string `json:"_id"`
	UserID string `json:"userId"`
	Stop   bool   `json:"stop"`
}
