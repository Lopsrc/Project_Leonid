package user

import "github.com/jackc/pgtype"

type User struct {
	ID int
	Name string
	Sex string
	Birthdate pgtype.Date
	Age int
	Weight int
}