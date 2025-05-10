package routes

import (
	"github.com/gin-gonic/gin"
	database "github.com/shun198/golang-clean-architecture/internal/infrastructures/databases"
	middleware "github.com/shun198/golang-clean-architecture/internal/infrastructures/middlewares"
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
	privateRoutes := r.Group(apiBase)
	privateRoutes.Use(middleware.AuthMiddleware())
	setupPublicRoutes(publicRoutes)
	setupPrivateRoutes(privateRoutes)
}

func setupPublicRoutes(publicRoutes *gin.RouterGroup) {
	setupHealthRoutes(publicRoutes)
	setupLoginRoutes(publicRoutes)
}

func setupPrivateRoutes(privateRoutes *gin.RouterGroup) {
	setupUserPrivateRoutes(privateRoutes)
}

func setupLoginRoutes(publicRoutes *gin.RouterGroup) {
	login := publicRoutes.Group("/login")
	loginRepo := repository.NewLoginRepository(database.DB)
	loginUseCase := usecase.NewLoginUseCase(loginRepo)
	loginHandler := handlers.NewLoginHandler(loginUseCase)
	login.POST("", loginHandler.Login)
}

func setupUserPrivateRoutes(privateRoutes *gin.RouterGroup) {
	users := privateRoutes.Group("/users")
	userRepository := repository.NewUserRepository(database.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(userUsecase)
	users.GET("", userHandler.GetUsers)
	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
}
