package main

import (
	"article/internal/config"
	"article/internal/handler"
	"article/internal/repository"
	"article/internal/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// --- Set Gin release mode (hilangkan warning debug) ---
	gin.SetMode(gin.ReleaseMode)

	// --- Init DB ---
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	// --- Setup repository, service, handler ---
	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)
	h := handler.NewPostHandler(svc)

	// --- Create Gin router ---
	r := gin.New() // Pakai New() agar Logger & Recovery bisa custom

	// --- Tambahkan middleware ---
	r.Use(gin.Recovery()) // error recovery
	r.Use(gin.Logger())   // logging

	// --- Tambahkan CORS ---
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // sesuaikan dengan domain front-end
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// --- Routes ---
	article := r.Group("/article")
	{
		article.POST("", h.Create)       // Create article
		article.GET("", h.GetAll)        // Get all articles
		article.GET("/:id", h.GetByID)   // Get article by ID
		article.PUT("/:id", h.Update)    // Update article (full)
		article.PATCH("/:id", h.Update)  // Update article (partial)
		article.DELETE("/:id", h.Delete) // Delete article
	}

	// --- Run server ---
	log.Println("Starting server on :8080 🚀")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}