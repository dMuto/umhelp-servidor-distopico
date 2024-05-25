package model

import "time"

type Currency struct {
	IDCurrency int64        `db: id_currency`
	Code       CurrencyCode `db:"code"`
	Symbol     string       `db:"symbol"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
}

type CurrencyCode string

const (
	CurrencyBRL CurrencyCode = "BRL"
)
