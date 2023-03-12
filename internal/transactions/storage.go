package transactions

import (
	"github.com/Chinedu-E/TradeTrail-go/internal/portfolios"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId   int
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
	db.AutoMigrate(&Transaction{})
	return &TransactionStorage{
		db: db,
	}
}

func (s *TransactionStorage) GetUserBalance(portfolioId int) (float64, error) {
	// Get the current virtual USD balance of the user's portfolio
	var balance float64
	if err := s.db.Table("portfolios").Select("balance").Where("id = ?", portfolioId).First(&balance).Error; err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *TransactionStorage) UpdateUserBalance(portfolioId int, balance float64) error {
	err := s.db.UpdateColumn("balance", balance).Where("id = ?", portfolioId).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionStorage) GetSecurityPrice(symbol string) (float64, error) {
	var price float64
	if err := s.db.Table("latest_prices").Where("symbol = ?", symbol).Select("price").Limit(1).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}

func (s *TransactionStorage) GetPortfolioTransactions(portfolioId int) ([]*Transaction, error) {
	var transactions []*Transaction
	err := s.db.Table("transactions").Where("source_id = ?", portfolioId).Scan(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionStorage) GetUserTransactions(userId int) ([]*Transaction, error) {
	var transactions []*Transaction
	err := s.db.Table("transactions").Where("user_id = ?", userId).Scan(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionStorage) GetUserShares(portfolioId int, symbol string) float64 {
	var shares float64
	err := s.db.Table("user_portfolios").Where("source_id = ? AND symbol = ?", portfolioId, symbol).Select("shares").Scan(&shares).Error
	if err != nil {
		shares = 0.0
	}
	return shares
}

func (s *TransactionStorage) UpdateUserPortfolio(portfolioId int, symbol string, shares float64) error {
	err := s.db.Table("user_portfolios").Where("source_id = ? AND symbol = ?", portfolioId, symbol).Update("shares", shares).Error
	return err
}

func (s *TransactionStorage) NewUserSecurity(portfolioId int, symbol string, shares float64) error {
	err := s.db.Table("user_portfolios").Create(&portfolios.UserPortfolio{PortfolioId: portfolioId, Symbol: symbol, Shares: shares}).Error
	return err
}

func (s *TransactionStorage) DeleteUserSecurity(portfolioId int, symbol string, shares float64) error {
	err := s.db.Table("user_portfolios").Delete(&portfolios.UserPortfolio{PortfolioId: portfolioId, Symbol: symbol, Shares: shares}).Error
	return err
}
