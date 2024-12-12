package models

type User struct {
	Id       int `db:"id"`
	Username string
	Password string
	Points   int
}

type UserInfo struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
}

type UserFullInfo struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}
