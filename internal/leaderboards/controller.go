package leaderboards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeaderBoardController struct {
	storage *LeaderBoardStorage
}

func NewLeaderBoardController(storage *LeaderBoardStorage) *LeaderBoardController {
	return &LeaderBoardController{
		storage: storage,
	}
}

func (t *LeaderBoardController) getLeaderBoard(c *gin.Context) {
	limit := 50
	leaderboard, err := t.storage.GetLeaderBoard(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, leaderboard)
}

func (t *LeaderBoardController) addToLeaderBoard(c *gin.Context) {
	var leaderboard *LeaderBoard
	if err := c.ShouldBind(&leaderboard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err := t.storage.AddToLeaderBoard(leaderboard)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, leaderboard)
}
