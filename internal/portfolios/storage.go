package portfolios

import (
	"time"

	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	UserId          int
	Type            string
	SecuritiesType  string
	StartingBalance float64
	Balance         float64
	Value           float64
	AutoManaged     bool
}

type PortfolioValues struct {
	gorm.Model
	PortfolioId int
	Date        time.Time
	Value       float64
}

type UserPortfolio struct {
	gorm.Model
	PortfolioId int
	Symbol      string `gorm:"not null"`
	Shares      float64
}

type PortfolioStorage struct {
	db *gorm.DB
}

func NewPortfolioStorage(db *gorm.DB) *PortfolioStorage {
	db.AutoMigrate(&Portfolio{})
	db.AutoMigrate(&PortfolioValues{})
	db.AutoMigrate(&UserPortfolio{})
	return &PortfolioStorage{
		db: db,
	}
}

func (s *PortfolioStorage) GetPortfolio(id int) (*Portfolio, error) {
	// Get the portfolio from the "Portfolios" table
	var portfolio *Portfolio
	if err := s.db.Where("id = ?", id).First(&portfolio).Error; err != nil {
		return nil, err
	}

	return portfolio, nil
}

func (s *PortfolioStorage) GetAllPortfolios() ([]*Portfolio, error) {
	// Get the portfolio from the "Portfolios" table
	var portfolios []*Portfolio
	if err := s.db.Find(&portfolios).Error; err != nil {
		return nil, err
	}

	return portfolios, nil
}

func (s *PortfolioStorage) GetAllUserPortfolios(id int) ([]*Portfolio, error) {
	// Get the portfolio from the "Portfolios" table
	var portfolios []*Portfolio

	// Get balance from the "Portfolios" table
	err := s.db.Table("portfolios").Where("user_id = ?", id).Scan(&portfolios).Error
	if err != nil {

		return nil, err
	}

	return portfolios, nil
}

func (s *PortfolioStorage) GetPortfolioHoldings(id int) ([]*UserPortfolio, error) {
	var holdings []*UserPortfolio
	if err := s.db.Table("user_portfolios").Where("portfolio_id = ?", id).Scan(&holdings).Error; err != nil {
		return nil, err
	}
	return holdings, nil
}

func (s *PortfolioStorage) createPortfolio(portfolio *Portfolio) error {
	if err := s.db.Create(&portfolio).Error; err != nil {

		return err
	}
	return nil
}
