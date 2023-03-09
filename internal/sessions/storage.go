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

func NewSessionStorage(db *gorm.DB) *SessionStorage {
	return &SessionStorage{
		db: db,
	}
}
func (s *SessionStorage) GetSession(id int) (*Session, error) {
	var session *Session
	err := s.db.First(&session, id).Error

	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionStorage) Create(session *Session) (*Session, error) {

	if err := s.db.Create(&session).Error; err != nil {

		return nil, err
	}
	return session, nil
}
