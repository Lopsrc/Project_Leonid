package user

import "github.com/jackc/pgtype"

type GetUser struct {
	ID int
}

type UpdateUser struct {
	Id int64
	Name string
	Sex string
	Birthdate pgtype.Date
	Age int
	Weight int
}


