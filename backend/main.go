package main

import (
	"log"
	"studybuddy/handlers"
	"studybuddy/middleware"
	"studybuddy/storage"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load existing users from file
	storage.LoadUsers()

	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve static files (frontend)
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/login", "./static/login.html")
	r.StaticFile("/register", "./static/register.html")
	r.StaticFile("/forgot-password", "./static/forgot-password.html")

	// Public authentication routes
	auth := r.Group("/auth")
	{
		auth.POST("/login", handlers.HandleLogin)
		auth.POST("/register", handlers.HandleRegister)
	}

	// Protected API routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/data", handlers.HandleGetData)
		api.POST("/data", handlers.HandleSaveData)
		api.DELETE("/events/:id", handlers.HandleDeleteEvent)
	}

	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}
