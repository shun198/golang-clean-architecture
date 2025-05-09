package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	"github.com/shun198/golang-clean-architecture/internal/presentation/responses"
)

type IUserUsecase interface {
	CreateUser(req requests.CreateUserRequest) (*models.User, error)
}

type UserHandler struct {
	userUsecase IUserUsecase
}

func NewUserHandler(userUsecase IUserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "validation error",
		})
	}
	user, err := h.userUsecase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot craete user",
		})
	}

	c.JSON(http.StatusCreated, responses.ListUsersResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		IsActive: user.IsActive,
	})
}
