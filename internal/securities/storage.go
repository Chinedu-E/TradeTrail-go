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
