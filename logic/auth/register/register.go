package register

import (
	"net/http"

	"github.com/Belxz777/backgo/common/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type registration struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	City     string `json:"city"`
}

func (h handler) register(c *gin.Context) {
	body := registration{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Email = body.Email
	user.Username = body.Username
	user.City = body.City
	user.ID = uuid.New().String()
	user.AvatarURL = ""        // Default empty avatar URL
	user.Bio = ""              // Default empty bio
	user.TotalWorkouts = 0     // Initialize workout count
	user.TotalDistanceKm = 0.0 // Initialize distance
	user.TotalCalories = 0     // Initialize calories
	user.IsMetric = true       // Default to metric system
	user.IsPublic = true       // Default to public profile

	result := h.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &user)
}
