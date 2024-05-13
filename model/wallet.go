package model

import "time"


type Wallet struct{
	ID 			int64 		`db:"id_wallet"`
	IDOwner 	int64 		`db:"id_owner"`
	Alias 		string 		`db:"alias"`
	IDCurrency 	int64 		`db:"id_currency"`
	Balancer 	int64 		`db:"balancer"`
	CreatedAt 	time.Time 	`db:"created_at"`
}