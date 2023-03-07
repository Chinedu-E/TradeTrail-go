package users

import (
	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(app *gin.Engine, controller *UserController) {
	users := app.Group("/users")
	users.GET("/all", controller.getAllUsers)
	users.GET("/user", controller.getUser)
	users.POST("/login", controller.login)
	users.POST("/register", controller.register)
}
