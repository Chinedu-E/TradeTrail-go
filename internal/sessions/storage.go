package sessions

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	HostId          int
	Duration        int
	Symbol          string
	StartingBalance float64
	NumParticipants int
	AgainstAgent    bool
}

type SessionParticipants struct {
	gorm.Model
	SessionId    int
	UserId       int
	StartBalance float64
	NumTrades    int
	EndBalance   float64
	NetShares    float64
	Profit       float64
	Agent        bool
}

type SessionStorage struct {
	db *gorm.DB
}
