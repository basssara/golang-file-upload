package jwtToken

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Username  string `json:"username"`
	ID        string `json:"id"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

var secretKey = []byte("secret")

func CreateToken(u UserClaims) (string, error) {
	tokenTTL, _ := strconv.Atoi("1800") //os.GetEnv("TOKEN_TTL")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        u.ID,
		"username":  u.Username,
		"createdAt": u.CreatedAt,
		"role":      u.Role,
		// "exp":       time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err

	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	roles := []string{"user", "admin"}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token")

	}

	authorized := false

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		role, err := claims["role"].(string)

		if !err && role == "" {
			return errors.New("empty role")
		}

		for _, roleRange := range roles {
			if roleRange == role {
				authorized = true
				break
			}
		}

		if !authorized {
			return errors.New("unauthorized")

		}
	}

	return nil

}

// func ValidateJWT(ctx *gin.Context) error {
// 	token, err := getToken(ctx)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return err
// 	}

// 	_, ok := token.Claims.(jwt.MapClaims)

// 	if !ok && !token.Valid {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return errors.New("invalid token")
// 	}

// 	return nil
// }

// func ValidateAdminRoleJWT(ctx *gin.Context) error {
// 	token, err := getToken(ctx)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	role := claims["role"].(string)

// 	if !ok && !token.Valid && role != "admin" {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return errors.New("invalid author token provided")
// 	}
// 	return nil
// }

// func ValidateUserRoleJWT(ctx *gin.Context) error {
// 	token, err := getToken(ctx)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	role := claims["role"].(string)

// 	if !ok && !token.Valid && role != "user" {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return errors.New("invalid author token provided")
// 	}
// 	return nil
// }

// func getToken(ctx *gin.Context) (*jwt.Token, error) {
// 	tokenString := getTokenFromRequest(ctx)

// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}
// 		return secretKey, nil
// 	})

// 	return token, err
// }

// func getTokenFromRequest(ctx *gin.Context) string {
// 	bearerToken := ctx.Request.Header.Get("Authorization")

// 	splitToken := strings.Split(bearerToken, " ")

// 	if len(splitToken) == 2 {
// 		return splitToken[1]
// 	}
// 	return ""
// }
