package securities

import "github.com/gin-gonic/gin"

func AddSecuritiesRoutes(app *gin.Engine, controller *SecurityController) {
	prices := app.Group("/prices")
	prices.GET("/info", controller.getSecurityInfo)
	prices.GET("/latest", controller.getLatestPrice)
	prices.GET("/latest/:limit", controller.getLatestPrices)
	prices.GET("/history", controller.getSymbolHistory)
}
