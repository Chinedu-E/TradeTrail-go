package consumers

import (
	"encoding/json"
	"log"

	"github.com/Chinedu-E/TradeTrail-go/internal/securities"
	"gorm.io/gorm"
)

func SymbolUpdates(db *gorm.DB) {
	msgs := ConsumeChannel("symbols")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var securities []*securities.SecuritiesInfo
			json.Unmarshal(d.Body, &securities)
			if err := db.Create(&securities).Error; err != nil {
				log.Printf("There was an error: %s", err.Error())
			}

		}
	}()

	log.Printf(" [*] Waiting for symbol updates. To exit press CTRL+C")
	<-forever
}
