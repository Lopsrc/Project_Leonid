package model

type AuthData struct {
	Id int					`json:"id"`
	Login string			`json:"login"`
	State bool				`json:"state"`
	Access_token string		`json:"access_token"`
	Refresh_token string	`json:"refresh_token"`
}
type UserData struct {
	Id 			int			`json:""`
	Name 		string		`json:""`
	Sex 		string		`json:""`
	Bithdate 	string		`json:""`
	Weight 		int			`json:""`
}
