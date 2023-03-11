package leaderboards

import "github.com/gin-gonic/gin"

func AddLeaderboardRoutes(app *gin.Engine, controller *LeaderBoardController) {
	leaderboard := app.Group("/leaderboard")
	leaderboard.GET("/", controller.getLeaderBoard)
	leaderboard.POST("/", controller.addToLeaderBoard)
}
