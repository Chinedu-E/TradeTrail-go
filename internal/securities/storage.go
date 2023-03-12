package securities

import (
	"gorm.io/gorm"
)

type LatestPrices struct {
	gorm.Model
	Symbol               string  `gorm:"unique;not null" json:"Symbol"`
	Price                float64 `json:"Price"`
	The1YTargetEst       float64 `json:"1y Target Est"`
	The52WeekRange       string  `json:"52 Week Range"`
	Ask                  string  `json:"Ask"`
	AvgVolume            int64   `json:"Avg. Volume"`
	Beta5YMonthly        float64 `json:"Beta (5Y Monthly)"`
	Bid                  string  `json:"Bid"`
	DaySRange            string  `json:"Day's Range"`
	EpsTtm               float64 `json:"EPS (TTM)"`
	EarningsDate         string  `json:"Earnings Date"`
	ExDividendDate       string  `json:"Ex-Dividend Date"`
	ForwardDividendYield string  `json:"Forward Dividend & Yield"`
	MarketCap            string  `json:"Market Cap"`
	Open                 float64 `json:"Open"`
	PERatioTTM           float64 `json:"PE Ratio (TTM)"`
	PreviousClose        float64 `json:"Previous Close"`
	QuotePrice           float64 `json:"Quote Price"`
	Volume               float64 `json:"Volume"`
}

type SecuritiesInfo struct {
	gorm.Model
	Symbol     string `gorm:"unique;not null"`
	Name       string
	Sector     string
	SubSector  string
	HQlocation string
	DateAdded  string
	Founded    string
}

type SecurityPrices struct {
	gorm.Model
	Date   string
	Symbol string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

type SecurityStorage struct {
	db *gorm.DB
}

func NewTransactionStorage(db *gorm.DB) *SecurityStorage {
	db.AutoMigrate(&LatestPrices{})
	db.AutoMigrate(&SecuritiesInfo{})
	db.AutoMigrate(&SecurityPrices{})
	return &SecurityStorage{
		db: db,
	}
}

func (s *SecurityStorage) GetSecuritiesInfo(symbol string) (*SecuritiesInfo, error) {
	var info *SecuritiesInfo
	if err := s.db.Table("securities_info").Where("symbol = ?", symbol).Scan(&info).Error; err != nil {
		return nil, err
	}
	return info, nil
}

func (s *SecurityStorage) GetLiveInfo(symbol string) (*LatestPrices, error) {
	var info *LatestPrices
	if err := s.db.Table("latest_prices").Where("symbol = ?", symbol).Scan(&info).Error; err != nil {
		return nil, err
	}
	return info, nil
}

func (s *SecurityStorage) GetLatestPrices(limit int) ([]*LatestPriceResponse, error) {
	var prices []*LatestPriceResponse
	if err := s.db.Table("latest_prices").Order("volume DESC").Select("securities.name, latest_prices.symbol, latest_prices.price, latest_prices.previous_close").Joins("join securities on latest_prices.symbol = securities.symbol").Limit(limit).Scan(&prices).Error; err != nil {
		return nil, err
	}
	return prices, nil
}

func (s *SecurityStorage) GetLatestPrice(symbol string) (float64, error) {
	var price float64
	if err := s.db.Table("latest_prices").Where("symbol = ?", symbol).Select("quote_price").Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}

func (s *SecurityStorage) GetSymbolHistory(symbol string, numDays int) ([]*SecurityPrices, error) {
	var history []*SecurityPrices
	if err := s.db.Table("security_prices").Order("date DESC").Where("symbol = ?", symbol).Limit(numDays).Scan(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}
