package watchlists

import (
	"gorm.io/gorm"
)

type WatchList struct {
	gorm.Model
	UserId int
	Symbol string
}

type WatchListStorage struct {
	db *gorm.DB
}

func NewWatchListStorage(db *gorm.DB) *WatchListStorage {
	return &WatchListStorage{
		db: db,
	}
}
func (s *WatchListStorage) GetAll() ([]*WatchList, error) {
	var list []*WatchList
	err := s.db.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *WatchListStorage) GetWatchlist(id int) (*WatchList, error) {
	var watchlist *WatchList
	err := s.db.First(&watchlist, id).Error

	if err != nil {
		return nil, err
	}
	return watchlist, nil
}

func (s *WatchListStorage) InWatchList(userId int, symbol string) (*WatchList, error) {
	var watchList *WatchList
	if err := s.db.Where("user_id = ? AND symbol = ?", userId, symbol).First(watchList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return watchList, nil
		} else {
			return nil, err
		}
	}
	return watchList, nil
}

func (s *WatchListStorage) AddToWatchList(userId int, symbol string) (*WatchList, error) {
	watchlist, err := s.InWatchList(userId, symbol)
	if err != nil {
		return nil, err
	}

	if err := s.db.Create(&watchlist).Error; err != nil {
		return nil, err
	}
	return watchlist, nil
}

func (s *WatchListStorage) RemoveFromWatchList(userId int, symbol string) error {
	watchlist, err := s.InWatchList(userId, symbol)
	if err != nil {
		return err
	}
	if err := s.db.Table("watch_lists").Where("user_id = ? AND symbol = ?", watchlist.UserId, watchlist.Symbol).Delete(&watchlist).Error; err != nil {
		return err
	}

	return nil
}

func (s *WatchListStorage) GetUserWatchList(userId int) ([]string, error) {
	var watchlist []string
	err := s.db.Table("watch_lists").Select("symbol").Where("user_id = ? AND deleted_at is NULL", userId).Scan(&watchlist).Error
	if err != nil {
		return nil, err
	}

	return watchlist, nil
}
