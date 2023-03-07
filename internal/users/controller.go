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

	err := t.storage.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (t *UserController) login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	exists := t.storage.CheckExistingEmail(email)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username",
		})
	}
	// Verify password
	err := t.storage.CheckPassword(email, password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// Generate jwt
	tokenString, err := generateJWT(email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return JWT token to client
	c.JSON(200, gin.H{"token": tokenString})
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
