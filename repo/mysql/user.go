package mysql

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	cli *sqlx.DB
}

func (r *User) InsertUser(ctx context.Context, user *model.User) error {
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

	_, err := r.cli.ExecContext(ctx, query, user.Name, user.LastName, user.Document, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
