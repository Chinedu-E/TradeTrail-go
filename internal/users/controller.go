package users

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

func (t *UserController) register(c *gin.Context) error {
	return nil
}

func (t *UserController) login(c *gin.Context) error {
	return nil
}

func (t *UserController) getUser(c *gin.Context) error { return nil }

func (t *UserController) getAllUsers(c *gin.Context) error { return nil }
