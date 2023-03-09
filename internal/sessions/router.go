package sessions

import "github.com/gin-gonic/gin"

func AddSessionRoutes(app *gin.Engine, controller *SessionController) {
	session := app.Group("/session")
	session.GET("/", controller.getSession)
	session.POST("/", controller.createSession)

}
