package controllers

import (
	"net/http"
	"os"
	"time"
	"v1/src/initializers"
	"v1/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Pega o email/senha do body da requisição
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body request",
		})

		return
	}

	// Faz o hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})

		return
	}

	// Cria o usuário
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"senha": user.Password,
	})
}

func Login(c *gin.Context) {
	// Pega o email/senha do body da requisição
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body request",
		})

		return
	}

	// Procura o usuário no banco de dados
	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password incorrect",
		})

		return
	}

	// Compara a senha informada com a senha hasheada armazenada
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password incorrect",
		})

		return
	}

	// Gera um token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":    user.ID,
		"expiration": time.Now().Add(time.Hour * 24 * 30).Unix(), // Expira em 30 dias
	})

	// Realiza o login e pega o token codificado, usando a chave secreta do JWT
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Armazena o token nos cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) //	SECURE FALSE PQ POR ENQUANTO RODA LOCAL, HTTP ONLY TRUE PQ ESTA NO LOCALHOST

	// Success
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
