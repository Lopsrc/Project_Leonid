package userdata

type CreateUserData struct{
	Name string		`json:"name"`
	Sex string		`json:"sex"`
	Birthdate string `json:"birthdate"`
	Weight int `json:"weight"`
}