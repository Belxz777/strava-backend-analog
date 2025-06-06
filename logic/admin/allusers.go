package admin

import (
	"net/http"

	"github.com/Belxz777/backgo/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) getAllUsers(c *gin.Context) {
	var users []models.User
	if result := h.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &users)
}
