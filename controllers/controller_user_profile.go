package controllers

import (
	"net/http"
	"strconv"

	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/app"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/database"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/helpers"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// A function to registered a user account
func RegisterUserAccount(context *gin.Context) {
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error, can't connect to the database!",
		})
		return
	}

	var newUser app.UserProfileData
	if error := context.BindJSON(&newUser); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	insertUser := models.UserProfile{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: helpers.EncryptPassword(newUser.Password),
	}

	_, error = govalidator.ValidateStruct(insertUser)

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": error.Error(),
		})
		return
	}

	var checkEmail models.UserProfile
	conn.Where("email = ?", newUser.Email).First(&checkEmail)
	var checkUsername models.UserProfile
	conn.Where("username = ?", newUser.Username).First(&checkUsername)

	if checkEmail.Email != "" || checkUsername.Username != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Username or Email has already exist!",
		})
		return
	}

	result := conn.Create(&insertUser)

	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Username or Email has already exist!",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully registered!",
	})
}

func LoginUserAccount(context *gin.Context) {
	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error, can't connect to the database!",
		})
		return
	}

	var user models.UserProfile
	email := context.Query("email")
	password := context.Query("password")
	err := conn.Where("email = ?", email).First(&user).Error

	if err != nil || !helpers.CheckPassword(password, user.Password) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	token, err := helpers.GenerateToken(user)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error, can't generate the token!",
		})
		return
	}

	context.SetCookie("Authorization", token, 3600, "", "", true, true)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully logged in!",
	})
}

// A function to logout from a user account
func LogoutUserAccount(context *gin.Context) {

	_, err := context.Cookie("Authorization")

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "Unauthorized Access!",
		})
		return
	}

	context.SetCookie("Authorization", "", -1, "", "", true, true)

	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully Logout!",
	})
}

// A function to update the user account
func UpdateUserAccount(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error, can't connect to the database!",
		})
		return
	}

	updateID := context.Param("id")

	var newUser app.UserProfileData
	if error := context.ShouldBindJSON(&newUser); error != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"status":  404,
			"message": "Invalid credentials!",
		})
		return
	}

	userData := context.MustGet("user").(models.UserProfile)

	if strconv.Itoa(userData.ID) != updateID {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "You dont have access to update this user account!",
		})
		return
	}

	var user models.UserProfile
	conn.Where("id = ?", updateID).First(&user)
	var checkEmail models.UserProfile
	conn.Where("email = ?", newUser.Email).First(&checkEmail)
	var checkUsername models.UserProfile
	conn.Where("username = ?", newUser.Username).First(&checkUsername)

	if checkEmail.Email != "" || checkUsername.Username != "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Username or Email has already exist!",
		})
		return
	}

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = helpers.EncryptPassword(newUser.Password)
	_, error = govalidator.ValidateStruct(user)

	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": error.Error(),
		})
		return
	}

	conn.Save(&user)
	context.SetCookie("Authorization", "", -1, "", "", true, true)
	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully updated the user account, please login to the related user account if you want continue updating it!",
	})

}

// A function to delete a user account
func DeleteUserAccount(context *gin.Context) {

	conn, error := database.Connect()

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error, can't connect to the database!",
		})
		return
	}

	deleteID := context.Param("id")
	var user models.UserProfile
	conn.Where("id = ?", deleteID).First(&user)
	userData := context.MustGet("user").(models.UserProfile)

	if strconv.Itoa(userData.ID) != deleteID {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "You dont have access to delete this user account, please login to the related user account if you want continue!",
		})
		return
	}

	conn.Delete(&user)
	context.SetCookie("Authorization", "", -1, "", "", true, true)
	context.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Successfully deleted this user account, please register or login again to continue!",
	})

}
