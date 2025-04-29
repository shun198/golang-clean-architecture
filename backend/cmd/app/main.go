package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("サーバの起動に失敗しました: %v", err)
	}
}
