package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"file-system/pkg/helpers"
	jwtToken "file-system/pkg/jwt"
)

func LoginHandler(ctx *gin.Context) {

	var u User

	if err := ctx.BindJSON(&u); err != nil {
		helpers.ErrorHelper(err)
	}

	userClaims := jwtToken.UserClaims{
		ID:       u.Email,
		Username: u.Password,
	}

	json.NewDecoder(ctx.Request.Body).Decode(&u)

	if u.Email == "JoeB@gmail.com" && u.Password == "swwwW112" {
		tokenString, err := jwtToken.CreateToken(userClaims)

		if err != nil {
			helpers.ErrorHelper(err)
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"token": tokenString})

		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}

func ProtectedHandler(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")
	tokenString := ctx.Request.Header.Get("Authorization")

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := jwtToken.VerifyToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "protected"})
	ctx.Next()
}
