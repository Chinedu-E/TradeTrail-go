package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	db.AutoMigrate(&User{})
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) CheckExistingEmail(email string) bool {
	var existingUser User
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err != nil {
		// An error occurred when querying the database, return an error message
		return true
	} else {
		return false
	}

}

func (s *UserStorage) CheckPassword(email string, password string) error {
	var user User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func (s *UserStorage) CreateUser(username, email, password string) (err error) {

	user := &User{Username: username, Email: email, Password: password}

	if err := s.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserStorage) GetUser(userId int64) (*User, error) {
	var user *User
	if err := s.db.First(&user, userId).Error; err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

func (s *UserStorage) GetAllUsers() (*[]User, error) {
	var users *[]User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}

}
