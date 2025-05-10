package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	"github.com/shun198/golang-clean-architecture/internal/presentation/responses"
	usecase "github.com/shun198/golang-clean-architecture/internal/usecases"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) *UserHandler {
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
		return
	}
	user, err := h.userUsecase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "cannot craete user",
		})
		return
	}

	c.JSON(http.StatusCreated, responses.ListUsersResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		IsActive: user.IsActive,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "id is not a valid integer",
		})
		return
	}
	user, err := h.userUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(http.StatusCreated, responses.ListUsersResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		IsActive: user.IsActive,
	})
}
