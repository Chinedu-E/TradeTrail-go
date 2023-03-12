package consumers

import (
	"encoding/json"
	"log"

	"github.com/Chinedu-E/TradeTrail-go/internal/securities"
	"gorm.io/gorm"
)

func LivePrices(db *gorm.DB) {
	msgs := ConsumeChannel("live_prices")
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var securities securities.LatestPrices
			json.Unmarshal(d.Body, &securities)
			if err := db.Table("latest_prices").Where("symbol = ?", securities.Symbol).UpdateColumns(securities).Error; err != nil {
				log.Printf("There was an error: %s", err.Error())
			}
			go updateUsersPortfolioValue(db, securities.Symbol)
		}
	}()

	log.Printf(" [*] Waiting for symbol updates. To exit press CTRL+C")
	<-forever
}
