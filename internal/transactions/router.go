package transactions

import "github.com/gin-gonic/gin"

func AddTransactionRoutes(app *gin.Engine, controller *TransactionController) {
	transactions := app.Group("/transactions")
	transactions.GET("/portfolio/:id", controller.getPortfolioTransactions)
	transactions.GET("/user/:id", controller.getUserTransactions)

	transactions.POST("/buy", controller.buy)
	transactions.POST("/sell", controller.sell)
}
