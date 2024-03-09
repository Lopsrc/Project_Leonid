package auth

type User struct {
	ID int64
	Email string
	Passhash []byte
	IsDeleted bool
}