package auth

type CreateUser struct {
	Email string 
	Password string 
}

type LoginUser struct {
	Email string
	Password string
}

type UpdateUser struct {
	Id int64
	Passhash []byte
}

type DeleteUser struct {
	Id int64
}

type RecoverUser struct {
	Id int64
}

