package userdata

import (
	"github.com/jackc/pgtype"
)

type UserData struct{
	Id int	`json:"id"`
	Name string	`json:"name"`
	Sex string	`json:"sex"`
	Birthdate pgtype.Date 	`json:"birthdate"`
	Weight int	`json:"weight"`
}