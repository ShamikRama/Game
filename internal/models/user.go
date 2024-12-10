package models

type User struct {
	Id       int `db:"db"`
	Username string
	Password string
	Points   int
}
