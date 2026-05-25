package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *gin.Context) {
	// отримуємо заголовок Authorization з запиту
	authHeader := c.GetHeader("Authorization")

	// перевіряємо що заголовок є і починається з "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен відсутній"})
		c.Abort() // зупиняємо обробку запиту
		return
	}

	// витягуємо сам токен (прибираємо "Bearer " з початку)
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// парсимо та перевіряємо токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// перевіряємо що алгоритм підпису правильний
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невірний або прострочений токен"})
		c.Abort()
		return
	}

	// витягуємо дані з токена і кладемо в контекст
	claims := token.Claims.(jwt.MapClaims)
	c.Set("user_id", claims["user_id"])
	c.Set("username", claims["username"])

	// передаємо запит далі
	c.Next()
}
