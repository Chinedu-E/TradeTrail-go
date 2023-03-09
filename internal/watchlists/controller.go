package watchlists

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WatchListController struct {
	storage *WatchListStorage
}

func NewWatchListController(storage *WatchListStorage) *WatchListController {
	return &WatchListController{
		storage: storage,
	}
}

func (t *WatchListController) getWatchlist(c *gin.Context) {
	id := c.Query("id")
	watchId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	watchlist, err := t.storage.GetWatchlist(watchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, watchlist)
}

func (t *WatchListController) getWatchlists(c *gin.Context) {
	watchLists, err := t.storage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, watchLists)
}

func (t *WatchListController) getUserWatchList(c *gin.Context) {
	id := c.Query("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	watchlist, err := t.storage.GetUserWatchList(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, watchlist)
}

func (t *WatchListController) addToWatchList(c *gin.Context) {
	var watchList *WatchList
	if err := c.ShouldBind(&watchList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	_, err := t.storage.AddToWatchList(watchList.UserId, watchList.Symbol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, watchList)
}

func (t *WatchListController) removeFromWatchList(c *gin.Context) {
	var watchList *WatchList
	if err := c.ShouldBind(&watchList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err := t.storage.RemoveFromWatchList(watchList.UserId, watchList.Symbol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.Status(200)
}
