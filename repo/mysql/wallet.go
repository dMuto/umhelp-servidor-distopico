package mysql

import (
	"context"

	"github.com/dMuto/umhelp-servidor-distopico/model"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	cli *sqlx.DB
}

func (r *Wallet) InsertWallet(tx *sqlx.Tx, ctx context.Context, wallet *model.Wallet) error {
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
	exec := r.cli.ExecContext
	if tx != nil{
		exec = tx.ExecContext
	}

	_, err := exec(ctx, query, wallet.IDOwner, wallet.Alias, wallet.IDCurrency)
	if err != nil {
		return err
	}

	return nil
}

func (r *Wallet) VerifyAmount(tx *sqlx.Tx, ctx context.Context, idOwner int64) (*model.Wallet ,bool, error){

	var result model.Wallet
	
	query := `
	SELECT 
		id_wallet,
		balance 
	FROM 
		UMHELP.TAB_WALLET
	WHERE 
		id_owner = ?;
	`
	exec := r.cli.QueryxContext
	if tx != nil {
		exec = tx.QueryxContext
	}

	rows, err := exec(ctx, query, idOwner)
	if err != nil {
		return nil, false, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&result); err != nil {
			return  nil, false, err
		}
	} else {
		return nil, false, nil
	}

	if rows.Err() != nil {
		return  nil, false, rows.Err()
	}

	return &result, true, nil

}

func (r *Wallet) UpdateAmount(tx *sqlx.Tx, ctx context.Context, amountValue, idWallet int64 ) error{

	query := `
	UPDATE 
		UMHELP.TAB_WALLET
	SET
		balance = ?
	WHERE 
		id_wallet = ?;
	`
	exec := r.cli.ExecContext
	if tx != nil{
		exec = tx.ExecContext
	}

	_, err := exec(ctx, query, amountValue, idWallet)
	if err != nil {
		return err
	}

	return nil

}