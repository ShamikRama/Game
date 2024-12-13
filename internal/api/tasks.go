package api

import (
	"Game/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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
	// Получение ID пользователя
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "user not found")
		return
	}

	// Привязка JSON-запроса к структуре
	var ref models.RefInput
	if err := c.BindJSON(&ref); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	// Вызов сервиса
	err = h.services.EnterRefCode(userId, ref.Id)
	if err != nil {
		log.Printf("Error entering referral code: %v", err)

		if strings.Contains(err.Error(), "can not enter own ref_code") {
			newErrorResponse(c, http.StatusConflict, "can not enter own ref_code")
			return
		}

		if strings.Contains(err.Error(), "referrer by ID not found") {
			newErrorResponse(c, http.StatusNotFound, "referrer not found")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "server error")
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, map[string]interface{}{
		"entered referral code": ref.Id,
	})
}
