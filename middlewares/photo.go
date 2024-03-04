package middlewares

import (
	"net/http"

	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/database"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/models"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {

	return func(context *gin.Context) {

		photoId := context.Param("id")

		conn, error := database.Connect()

		if error != nil {
			context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"status":  500,
				"message": "Error, can't connect to the database!",
			})
			return
		}

		var photo models.UserPhoto
		conn.Where("id = ?", photoId).First(&photo)

		if photo.ID == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  404,
				"message": "Photo not found!",
			})
			return
		}

		userData := context.MustGet("user").(models.UserProfile)

		if photo.UserID != userData.ID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "You dont have access to this photo!",
			})
			return
		}

		context.Next()
	}
}
