package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) CompleteTaskTelegram(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}

	err = h.services.CompleteTaskTelegram(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"task": "telegram",
	})
}
