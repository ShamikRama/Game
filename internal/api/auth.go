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

	id, err := h.services.Create(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handlers) SignIn(c *gin.Context) {
	var input models.Login
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad request")
		return
	}
	token, err := h.services.Authorization.GenerateJwtToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
