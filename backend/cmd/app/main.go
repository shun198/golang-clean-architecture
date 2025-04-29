package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shun198/golang-clean-architecture/internal/infrastructures/database"
)

func main() {
	database.InitDB()
	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("サーバの起動に失敗しました: %v", err)
	}
}
