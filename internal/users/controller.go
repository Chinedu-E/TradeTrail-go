package users

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

func (t *UserController) register(c *gin.Context) {}

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
