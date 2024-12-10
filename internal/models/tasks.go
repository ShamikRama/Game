package models

type Tasks struct {
	Id       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Type     string `json:"type" db:"type"`
	Exp      int    `json:"exp" db:"exp"`
	Complete bool   `json:"complete" db:"complete"`
}
