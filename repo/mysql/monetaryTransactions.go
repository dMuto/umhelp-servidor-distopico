package mysql

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/jmoiron/sqlx"
)

type MonetaryTransaction struct {
	cli *sqlx.DB
}

func (r *MonetaryTransaction) NewMonetaryTransaction(tx *sqlx.Tx, ctx context.Context, mt model.MonetaryTransactions)  error {

	//var result model.Currency

	query := `
	INSERT INTO 
		UMHELP.TAB_MONETARY_TRANSACTIONS
	(
		id_origin,
		id_destination,
		amount_sent
	)
	VALUES
	(
		?,
		?,
		?
	);`

	exec := r.cli.ExecContext
	if tx != nil{
		exec = tx.ExecContext
	}

	_, err := exec(ctx, query, mt.IDOrigin, mt.IDDestination, mt.AmountSent)
	if err != nil {
		return err
	}


	return nil
}