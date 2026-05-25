package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"jwt-auth/database"
	"jwt-auth/models"
)

// Register - реєстрація нового користувача
func Register(c *gin.Context) {
	var input models.RegisterInput

	// перевіряємо що JSON правильний і всі поля є
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// хешуємо пароль перед збереженням (ніколи не зберігаємо пароль як є!)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Помилка хешування паролю"})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	// зберігаємо користувача в БД
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Користувач вже існує"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Користувача створено успішно!"})
}

// Login - логін та отримання JWT токена
func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// шукаємо користувача в БД
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невірний логін або пароль"})
		return
	}

	// порівнюємо пароль з хешем
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невірний логін або пароль"})
		return
	}

	// створюємо JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // токен діє 24 години
	})

	// підписуємо токен секретним ключем
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Помилка створення токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// GetProfile - захищений роут, повертає дані користувача
func GetProfile(c *gin.Context) {
	// отримуємо дані які middleware поклав у контекст
	username := c.MustGet("username").(string)
	userID := c.MustGet("user_id")

	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"username": username,
		"message":  "Це захищений роут!",
	})
}
