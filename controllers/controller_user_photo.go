package controllers

import (
	"net/http"

	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/app"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/database"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/models"
	"github.com/gin-gonic/gin"
)

// A function to create the user photo
func CreateUserPhoto(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	var newPhoto app.UserPhotoData
	if error := context.BindJSON(&newPhoto); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	userData := context.MustGet("user").(models.UserProfile)

	insertPhoto := models.UserPhoto{
		Title:    newPhoto.Title,
		Caption:  newPhoto.Caption,
		PhotoURL: newPhoto.PhotoURL,
		UserID:   userData.ID,
	}

	conn.Create(&insertPhoto)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully added new photo!",
	})
}

// A function to get the user photo
func GetUserPhoto(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	var photos []models.UserPhoto
	conn.Find(&photos)

	context.IndentedJSON(http.StatusOK, gin.H{
		"result":  photos,
		"status":  200,
		"message": "Success!",
	})
}

// A function to update the user photo
func UpdateUserPhoto(context *gin.Context) {

	updateID := context.Param("id")
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	var newPhoto app.UserPhotoData
	if error := context.BindJSON(&newPhoto); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	var photo models.UserPhoto
	conn.Where("id = ?", updateID).First(&photo)
	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoURL = newPhoto.PhotoURL
	conn.Save(&photo)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully updated the photo detail!",
	})
}

// A function to delete the user photo
func DeleteUserPhoto(context *gin.Context) {

	deleteID := context.Param("id")
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	var photo models.UserPhoto
	conn.Where("id = ?", deleteID).First(&photo)
	conn.Delete(&photo)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully deleted the old photo!",
	})
}
