package middlewares

import (
	"net/http"
	"time"

	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/database"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/helpers"
	"github.com/adityafd/task-5-pbi-fullstack-developer-adityafataha/models"
	"github.com/gin-gonic/gin"
)

func JwtCheck() gin.HandlerFunc {

	return func(context *gin.Context) {
		auth_token, err := context.Cookie("Authorization")

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "Unauthorized due to the Token not found!",
			})
			return
		}

		claims, err := helpers.ParseToken(auth_token)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "Unauthorized!",
			})
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  4000,
				"message": "Token Expired!",
			})
			return
		}

		var user models.UserProfile
		conn, err := database.Connect()

		conn.Where("email = ?", claims["email"]).First(&user)

		if user.ID == 0 || err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  4000,
				"message": "Token Invalid!",
			})
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
