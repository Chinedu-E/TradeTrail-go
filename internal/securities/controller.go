package securities

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LatestPriceResponse struct {
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	PreviousClose float64 `json:"previous_close"`
	Change        float64 `json:"change"`
}

type SecurityController struct {
	storage *SecurityStorage
}

func NewSecurityController(storage *SecurityStorage) *SecurityController {
	return &SecurityController{
		storage: storage,
	}
}

func (t *SecurityController) getLatestPrices(c *gin.Context) {
	limitStr := c.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid limit value")
	}
	prices, err := t.storage.GetLatestPrices(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	for i, price := range prices {
		prices[i].Change = (price.Price - price.PreviousClose) / price.PreviousClose
	}
	c.JSON(http.StatusOK, prices)
}

func (t *SecurityController) getSymbolHistory(c *gin.Context) {
	symbol := c.Query("symbol")
	n := c.Query("num_days")
	numDays, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid limit value")
	}

	history, err := t.storage.GetSymbolHistory(symbol, numDays)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, history)
}

func (t *SecurityController) getLatestPrice(c *gin.Context) {
	symbol := c.Query("symbol")
	price, err := t.storage.GetLatestPrice(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"price": price})
}

func (t *SecurityController) getSecurityInfo(c *gin.Context) {
	symbol := c.Query("symbol")

	info, err := t.storage.GetLiveInfo(symbol)
	if err != nil {
		log.Println(err)
	}

	extraInfo, err := t.storage.GetSecuritiesInfo(symbol)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"info": info, "extra": extraInfo})
}
