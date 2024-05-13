package model

type MonetaryTransactions struct {
	IDOrigin      int64 `db:"id_origin"`
	IDDestination int64 `db:"id_destination"`
	AmountSent    int64 `db:"amount_sent"`
}
