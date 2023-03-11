package portfolios

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PortfolioController struct {
	storage *PortfolioStorage
}

func NewPortfolioController(storage *PortfolioStorage) *PortfolioController {
	return &PortfolioController{
		storage: storage,
	}
}

func (t *PortfolioController) getPortfolio(c *gin.Context) {
	id := c.Query("id")
	portfolioId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	portfolio, err := t.storage.GetPortfolio(portfolioId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	holdings, _ := t.storage.GetPortfolioHoldings(portfolioId)

	c.JSON(http.StatusOK, gin.H{
		"portfolio": portfolio,
		"holdings":  holdings,
	})
}

func (t *PortfolioController) createPortfolio(c *gin.Context) {
	var portfolio *Portfolio
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := t.storage.createPortfolio(portfolio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, portfolio)
}

func (t *PortfolioController) getAllUserPortfolios(c *gin.Context) {
	id := c.Query("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	portfolios, err := t.storage.GetAllUserPortfolios(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, portfolios)
}
