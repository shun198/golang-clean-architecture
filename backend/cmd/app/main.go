package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	database "github.com/shun198/golang-clean-architecture/internal/infrastructures/databases"

	middleware "github.com/shun198/golang-clean-architecture/internal/infrastructures/middlewares"
	"github.com/shun198/golang-clean-architecture/internal/routes"
)

func main() {
	database.InitDB()
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Use(middleware.CORSConfig())
	routes.SetupRoutes(r)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("サーバの起動に失敗しました: %v", err)
	}
}
