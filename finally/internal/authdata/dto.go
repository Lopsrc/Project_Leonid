package authdata

type CreateAuthUserData struct{
	Login string	`json:"login"`
	Access_token string `json:"access_token"`
}