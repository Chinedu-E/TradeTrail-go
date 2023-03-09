package watchlists

import (
	"github.com/gin-gonic/gin"
)

func AddWatchListRoutes(app *gin.Engine, controller *WatchListController) {
	users := app.Group("/watchlist")
	users.GET("/", controller.getWatchlists)
	users.GET("/:id", controller.getWatchlist)
	users.GET("/user/:id", controller.getUserWatchList)
	users.POST("/add", controller.addToWatchList)
	users.POST("/remove", controller.removeFromWatchList)
}
