package watchlists

import (
	"github.com/gin-gonic/gin"
)

func AddWatchListRoutes(app *gin.Engine, controller *WatchListController) {
	watchlists := app.Group("/watchlist")
	watchlists.GET("/", controller.getWatchlists)
	watchlists.GET("/:id", controller.getWatchlist)
	watchlists.GET("/user/:id", controller.getUserWatchList)
	watchlists.POST("/add", controller.addToWatchList)
	watchlists.POST("/remove", controller.removeFromWatchList)
}
