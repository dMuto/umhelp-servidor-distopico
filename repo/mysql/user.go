package mysql

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) InsertUser(tx *sqlx.Tx, ctx context.Context, user *model.User) (userId int64, err error) {
	query := `
	INSERT INTO UMHELP.TAB_USER
	(
		name,
		last_name,
		document,
		email,
		password
	)
	VALUES
	(
		?,
		?,
		?,
		?,
		?
	);`
	
	exec := r.cli.ExecContext
	if tx != nil{
		exec = tx.ExecContext
	}

	result, err := exec(ctx, query, user.Name, user.LastName, user.Document, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
