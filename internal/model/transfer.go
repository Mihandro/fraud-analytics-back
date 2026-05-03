package model

import "time"

type Transfer struct {
	Id      int       `db:"id"`
	Amount  float64   `db:"amount"`
	Status  string    `db:"status"`
	SWallet float64   `db:"swallet"`
	DWallet float64   `db:"dwallet"`
	Date    time.Time `db:"date"`
}
