package authdata

type AuthData struct{
	Id int `json:"id"`
	Login string `json:"login"`
	State bool `json:"state"`
	Access_token string	`json:"access_token"`
	Refresh_token string	`json:"refresh_token"`
}