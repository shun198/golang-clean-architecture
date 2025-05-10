package routes

import (
	"github.com/gin-gonic/gin"
	database "github.com/shun198/golang-clean-architecture/internal/infrastructures/databases"
	"github.com/shun198/golang-clean-architecture/internal/presentation/handlers"
	repository "github.com/shun198/golang-clean-architecture/internal/repositories"
	usecase "github.com/shun198/golang-clean-architecture/internal/usecases"
)

func setupHealthRoutes(publicRoutes *gin.RouterGroup) {
	health := publicRoutes.Group("/health")
	{
		health.GET("", handlers.HealthCheck)
	}
}

func SetupRoutes(r *gin.Engine) {
	const apiBase = "/api"
	publicRoutes := r.Group(apiBase)

	setupPublicRoutes(publicRoutes)
}

func setupPublicRoutes(publicRoutes *gin.RouterGroup) {
	setupHealthRoutes(publicRoutes)
	users := publicRoutes.Group("/users")
	userRepository := repository.NewUserRepository(database.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(userUsecase)
	users.GET("/:id", userHandler.GetUser)
	users.POST("", userHandler.CreateUser)
}
