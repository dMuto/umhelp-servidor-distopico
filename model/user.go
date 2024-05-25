package model

type User struct {
	IdUser   int64  `db:"id_user"`
	Name     string `db:"name"`
	LastName string `db:"lastName"`
	Document string `db:"document"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
