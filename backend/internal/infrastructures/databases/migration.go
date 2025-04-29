package main

import (
	"fmt"
	"log"

	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	database "github.com/shun198/golang-clean-architecture/internal/infrastructures/databases"
)

func performMigration() error {
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("マイグレーションに失敗しました: %w", err)
	}
	return nil
}

func main() {
	database.InitDB()

	if err := performMigration(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("マイグレーションが正常に完了しました")
}
