package main

import (
	"jwt-auth/database"
	"jwt-auth/handlers"
	"jwt-auth/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("JWT_SECRET", "super-secret-key-1234")

	database.Connect()

	r := gin.Default()

	// публічні роути (без токена)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// захищені роути (потрібен токен)
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired)
	{
		auth.GET("/profile", handlers.GetProfile)
	}

	// запускаємо сервер на порту 8080
	r.Run(":8080")
}
