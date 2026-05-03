package repository

import (
	"fraud-analytics/internal/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetUsers() []model.User
	GetTransfers() []model.Transfer
}

type FraudTransfersRepository struct {
	DB *sqlx.DB
}

func NewFraudTransfersRepository() *FraudTransfersRepository {
	dsn := "postgres://:@localhost/postgres?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &FraudTransfersRepository{DB: db}
}

func (r *FraudTransfersRepository) GetTransfers() []model.Transfer {
	var transfers []model.Transfer

	err := r.DB.Select(&transfers, "SELECT * FROM transfers")
	if err != nil {
		log.Println(err)
	}

	return transfers
}

func (r *FraudTransfersRepository) GetUsers() []model.User {

	var users []model.User
	return users
}
