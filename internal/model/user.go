package model

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Lang     string `db:"lang"`
}
