package mysql

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r *Wallet) InsertWallet(ctx context.Context, wallet *model.Wallet) error {
	query := `
	INSERT INTO UMHELP.TAB_WALLET
	(
		id_owner,
		alias,
		id_currency
	)
	VALUES
	(
		?,
		?,
		?
	);
	`

	_, err := r.cli.ExecContext(ctx, query, wallet.IDOwner, wallet.Alias, wallet.IDCurrency)
	if err != nil {
		return err
	}

	return nil
}
