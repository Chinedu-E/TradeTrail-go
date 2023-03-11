package transactions

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SourceId int
	Symbol   string
	Shares   float64
	BoughtAt float64
	Type     string
	Source   string
}

type TransactionStorage struct {
	db *gorm.DB
}

func NewTransactionStorage(db *gorm.DB) *TransactionStorage {
	return &TransactionStorage{
		db: db,
	}
}

func CreateTransaction() {}
