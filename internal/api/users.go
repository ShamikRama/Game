package api

import (
	"Game/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) Status(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "not found user id")
		return
	}

	info, err := h.services.User.GetUserStatus(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, info)
}

func (h *Handlers) Leaders(c *gin.Context) {
	leaders, err := h.services.GetUsersLeaders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// TODO переместить логику в слой сервиса
	info := make([]models.UserInfo, 0, 5)
	for _, user := range leaders {
		info = append(info, models.UserInfo{
			Username: user.Username,
			Points:   user.Points,
		})
	}

	c.JSON(http.StatusOK, info)
}
