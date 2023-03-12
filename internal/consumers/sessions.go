package consumers

import (
	"encoding/json"
	"log"

	"github.com/Chinedu-E/TradeTrail-go/internal/leaderboards"
	"github.com/Chinedu-E/TradeTrail-go/internal/sessions"
	"gorm.io/gorm"
)

func SessionUpdates(db *gorm.DB) {
	msgs := ConsumeChannel("sessions")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var participant sessions.SessionParticipants
			json.Unmarshal(d.Body, &participant)
			if err := db.Create(&participant).Error; err != nil {
				log.Printf("There was an error: %s", err.Error())
			}
			addToLeaderBoard(db, participant)
		}
	}()

	log.Printf(" [*] Waiting for symbol updates. To exit press CTRL+C")
	<-forever
}

func addToLeaderBoard(db *gorm.DB, participant sessions.SessionParticipants) {
	var leaderBoard leaderboards.LeaderBoard

	err := db.Table("leaderboard").Where("user_id=?", participant.UserId).Scan(&leaderBoard).Error

	if err != nil {
		// User has participated in a session before
		var allTraded []string
		var sessionIds []int
		err := db.Table("session_participants").Where("user_id=?", participant.UserId).Select("session_id").Scan(&sessionIds).Error
		if err != nil {
			log.Println(err)
		}
		err = db.Table("sessions").Where("id IN", sessionIds).Select("symbol").Scan(&allTraded).Error
		if err != nil {
			log.Println(err)
		}
		leaderBoard.Profit += participant.Profit
		leaderBoard.MostTraded = mostCommonString(allTraded)
		db.UpdateColumns(leaderBoard)
	} else {
		var mostTraded string
		err := db.Table("sessions").Where("id=?", participant.SessionId).Select("symbol").Scan(&mostTraded).Error
		if err != nil {
			log.Println(err)
		}
		leaderBoard.Profit = participant.Profit
		leaderBoard.UserId = participant.UserId
		leaderBoard.MostTraded = mostTraded
		db.UpdateColumns(leaderBoard)
	}
}
