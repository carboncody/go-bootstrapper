package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/carboncody/go-bootstrapper/initializers"
	"github.com/carboncody/go-bootstrapper/models"
	"github.com/gin-gonic/gin"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userEmail := strings.TrimPrefix(authHeader, "Bearer ")

		// * Find the user from the email
		var user models.User
		result := initializers.DB.First(&user, "email = ?", fmt.Sprint(userEmail))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Email user in the header not found."})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
