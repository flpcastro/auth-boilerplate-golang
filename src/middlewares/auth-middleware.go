package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"v1/src/initializers"
	"v1/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Pega o token do cookie da requisição
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Valida se o token está correto
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verifica a data de expiração do mesmo
		if float64(time.Now().Unix()) > claims["expiration"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Acha o usuário com o subject do token
		var user models.User

		initializers.DB.First(&user, claims["subject"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Vincula com a requisição
		c.Set("user", user)

		// Passa o middleware
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
