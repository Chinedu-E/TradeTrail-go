package transactions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	storage *TransactionStorage
}

func NewTransactionController(storage *TransactionStorage) *TransactionController {
	return &TransactionController{
		storage: storage,
	}
}

func (t *TransactionController) buy(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Get the current virtual USD balance of the user's portfolio
	balance, err := t.storage.GetUserBalance(transaction.SourceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get the price of the security
	price, err := t.storage.GetSecurityPrice(transaction.Symbol)
	if err != nil {
		price = transaction.BoughtAt
	}

	totalCost := transaction.Shares * price

	// Check if the user has enough funds
	if totalCost > balance {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough funds"})
	}
	balance -= totalCost

	transaction.BoughtAt = price

	//Saving and Updating
	if err := t.storage.db.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := t.storage.UpdateUserBalance(transaction.SourceId, balance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Update the shares of the user's security portfolio
	shares := t.storage.GetUserShares(transaction.SourceId, transaction.Symbol)
	shares += transaction.Shares
	if shares == transaction.Shares { // check if the shares variable is equal to the new transaction shares
		// insert new record
		err = t.storage.NewUserSecurity(transaction.SourceId, transaction.Symbol, shares)
	} else {
		// update existing record
		err = t.storage.UpdateUserPortfolio(transaction.SourceId, transaction.Symbol, shares)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction completed successfully"})
}

func (t *TransactionController) sell(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Get the price of the security
	price, err := t.storage.GetSecurityPrice(transaction.Symbol)
	if err != nil {
		price = transaction.BoughtAt
	}

	totalCredit := transaction.Shares * price
	shares := t.storage.GetUserShares(transaction.SourceId, transaction.Symbol)
	if shares > transaction.Shares {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough shares to sell"})
	}
	shares += transaction.Shares
	if shares == 0.0 { // check if the shares variable is equal to the new transaction shares
		// delete old shares record
		err = t.storage.DeleteUserSecurity(transaction.SourceId, transaction.Symbol, shares)
	} else {
		// update existing record
		err = t.storage.UpdateUserPortfolio(transaction.SourceId, transaction.Symbol, shares)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := t.storage.db.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Get the current virtual USD balance of the user's portfolio
	balance, err := t.storage.GetUserBalance(transaction.SourceId)
	balance += totalCredit
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := t.storage.UpdateUserBalance(transaction.SourceId, balance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction completed successfully"})
}

func (t *TransactionController) getUserTransactions(c *gin.Context) {
	id := c.Query("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transactions, err := t.storage.GetUserTransactions(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, transactions)
}

func (t *TransactionController) getPortfolioTransactions(c *gin.Context) {
	id := c.Query("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transactions, err := t.storage.GetPortfolioTransactions(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, transactions)
}
