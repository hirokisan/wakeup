package domain

// Bed :
type Bed struct {
	ID     string `json:"_id"`
	UserID string `json:"userId"`
	Empty  bool   `json:"empty"`
}
