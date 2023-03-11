package portfolios

import "github.com/gin-gonic/gin"

func AddPortfolioRoutes(app *gin.Engine, controller *PortfolioController) {
	portfolio := app.Group("/portfolio")
	portfolio.GET("/:id", controller.getPortfolio)
	portfolio.GET("/user/:id", controller.getAllUserPortfolios)

	portfolio.POST("/", controller.createPortfolio)
}
