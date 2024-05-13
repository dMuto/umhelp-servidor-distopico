package model

type User struct {
	Name     string `db:"name"`
	LastName string `db:"lastName"`
	Document string `db:"document"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
