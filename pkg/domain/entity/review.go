package entity

type Review struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}
