package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

func (t *UserController) register(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	exists := t.storage.CheckExistingEmail(user.Email)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username already exists",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(password)
	t.storage.db.Create(&user)
	t.storage.db.Commit()

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (t *UserController) login(c *gin.Context) {

}

func (t *UserController) getUser(c *gin.Context) {}

func (t *UserController) getAllUsers(c *gin.Context) {}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   username,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"authorized": true,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
