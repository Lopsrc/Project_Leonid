package userdata

type UserData struct{
	Id int	`json:"id"`
	Name string	`json:"name"`
	Sex string	`json:"sex"`
	Birthdate string	`json:"birthdate"`
	Weight int	`json:"weight"`
}