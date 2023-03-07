package users

import (
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(app *gin.Engine, controller *UserController) {
	users := app.Group("/users")
	users.GET("/", controller.getAllUsers)
	users.GET("/:id", controller.getUser)
	users.POST("/login", controller.login)
	users.POST("/register", controller.register)
}
