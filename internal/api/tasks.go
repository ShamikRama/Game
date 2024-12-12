package api

import (
	"Game/internal/models"
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

func (h *Handlers) ReferralCode(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}

	var ref models.RefInput
	if err := c.BindJSON(&ref); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	err = h.services.EnterRefCode(userId, ref.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"entered referral code": ref.Id,
	})
}
