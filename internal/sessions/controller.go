package sessions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	storage *SessionStorage
}

func NewSessionController(storage *SessionStorage) *SessionController {
	return &SessionController{
		storage: storage,
	}
}

func (t *SessionController) getSession(c *gin.Context) {
	id := c.Query("id")
	sessionId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	session, err := t.storage.GetSession(sessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, session)
}

func (t *SessionController) createSession(c *gin.Context) {
	var session *Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if _, err := t.storage.Create(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, session)
}
