package main

import (
	"log"

	seed "github.com/shun198/golang-clean-architecture/internal/domains/seeds"
	database "github.com/shun198/golang-clean-architecture/internal/infrastructures/databases"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("パスワードのハッシュ化に失敗しました: %v", err)
	}
	return string(hashedPassword)
}

func main() {
	database.InitDB()

	users := seed.CreateUserLocalData()
	usersSuccessCount := 0

	for i := range users {
		users[i].Password = hashPassword(users[i].Password)
		if err := database.DB.Create(&users[i]).Error; err != nil {
			log.Printf("システムユーザの作成に失敗しました: %v", err)
			continue
		}
		usersSuccessCount++
	}

	if usersSuccessCount > 0 {
		log.Printf("システムユーザのテストデータを %d 件作成しました", usersSuccessCount)
	}

	todos := seed.CreateTodoLocalData()
	todosSuccessCount := 0

	for i := range todos {
		if err := database.DB.Create(&todos[i]).Error; err != nil {
			log.Printf("Todoの作成に失敗しました: %v", err)
			continue
		}
		todosSuccessCount++
	}

	if todosSuccessCount > 0 {
		log.Printf("Todoのテストデータを %d 件作成しました", todosSuccessCount)
	}
}
