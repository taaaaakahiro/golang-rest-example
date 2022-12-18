package entity

type Review struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}
