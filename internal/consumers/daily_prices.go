package consumers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Chinedu-E/TradeTrail-go/internal/portfolios"
	"github.com/Chinedu-E/TradeTrail-go/internal/securities"
	"gorm.io/gorm"
)

func DailyUpdates(db *gorm.DB) {
	msgs := ConsumeChannel("prices")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var security_price securities.SecurityPrices
			json.Unmarshal(d.Body, &security_price)

			if err := db.Create(&security_price).Error; err != nil {
				log.Printf("There was an error: %s", err.Error())
			}

			go updateUsersPortfolioValue(db, security_price.Symbol)
		}
	}()

	log.Printf(" [*] Waiting for symbol updates. To exit press CTRL+C")
	<-forever
}

func updateUsersPortfolioValue(db *gorm.DB, symbol string) {
	var portfolios_ []portfolios.UserPortfolio
	if err := db.Table("user_portfolios").Where("symbol=?", symbol).Scan(&portfolios_).Error; err != nil {
		return
	}

	for _, portfolio := range portfolios_ {
		var portfolioValue portfolios.PortfolioValues
		value := getPortfolioValues(db, portfolio.PortfolioId)
		if value == 0 {
			continue
		}
		portfolioValue.PortfolioId = int(portfolio.PortfolioId)
		portfolioValue.Value = value
		portfolioValue.Date = time.Now()

		db.Table("portfolios").Where("id = ?", portfolio.PortfolioId).UpdateColumn("value", value)
		db.Create(&portfolioValue)
	}
}

func getPortfolioValues(db *gorm.DB, portfolioId int) float64 {
	var holdings []portfolios.UserPortfolio
	var value float64
	if err := db.Table("user_portfolios").Where("portfolio_id = ?", portfolioId).Scan(&holdings).Error; err != nil {
		return 0
	}

	for _, holding := range holdings {
		var price float64
		if err := db.Table("latest_prices").Select("price").Where("symbol = ?", holding.Symbol).Scan(&price).Error; err != nil {
			log.Fatal(err)
		}
		pvalue := float64(price * holding.Shares)
		value += pvalue
	}

	return value
}
