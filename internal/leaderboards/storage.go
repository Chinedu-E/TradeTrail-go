package leaderboards

import "gorm.io/gorm"

type LeaderBoard struct {
	gorm.Model
	UserId     int `gorm:"unique"`
	Profit     float64
	MostTraded string
}

type LeaderBoardStorage struct {
	db *gorm.DB
}

func NewLeaderBoardStorage(db *gorm.DB) *LeaderBoardStorage {
	db.AutoMigrate(&LeaderBoard{})
	return &LeaderBoardStorage{
		db: db,
	}
}

func (s *LeaderBoardStorage) GetLeaderBoard(limit int) ([]*LeaderBoard, error) {
	var leaderboard []*LeaderBoard
	if err := s.db.Find(&leaderboard).Limit(limit).Error; err != nil {
		return nil, err
	} else {
		return leaderboard, nil
	}
}

func (s *LeaderBoardStorage) AddToLeaderBoard(leaderboard *LeaderBoard) error {
	if err := s.db.Create(&leaderboard).Error; err != nil {

		return err
	}
	return nil
}
