package api

import (
	"Game/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) SignUp(c *gin.Context) {
	var input models.Login
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	id, err := h.services.Create(user) // Передаем user
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
