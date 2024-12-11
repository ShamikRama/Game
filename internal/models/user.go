package models

type User struct {
	Id       int `db:"id"`
	Username string
	Password string
	Points   int
}
