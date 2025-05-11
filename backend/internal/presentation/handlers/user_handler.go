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

func (h *UserHandler) GetUsers(c *gin.Context) {
	var query requests.ListUsersQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if query.Limit == 0 {
		query.Limit = 20
	}

	results, err := h.userUsecase.GetAllUsers(query)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	users := make([]responses.UsersResponse, len(results.Users))
	for i, user := range results.Users {
		users[i] = responses.UsersResponse{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			Role:      user.Role,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			CreatedBy: user.CreatedBy,
			UpdatedBy: user.UpdatedBy,
		}
	}

	c.JSON(http.StatusOK, responses.ListUsersResponse{
		Count:  results.Total,
		Length: len(results.Users),
		Users:  users,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	auth_user_id := int(c.Keys["user_id"].(float64))
	user, err := h.userUsecase.CreateUser(req, auth_user_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.UsersResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
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
	c.JSON(http.StatusCreated, responses.UsersResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req requests.UpdateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
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
	auth_user_id := int(c.Keys["user_id"].(float64))
	updated_user, err := h.userUsecase.UpdateUser(req, user, auth_user_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, responses.UsersResponse{
		ID:        updated_user.ID,
		Email:     updated_user.Email,
		Username:  updated_user.Username,
		Role:      updated_user.Role,
		IsActive:  updated_user.IsActive,
		CreatedAt: updated_user.CreatedAt,
		UpdatedAt: updated_user.UpdatedAt,
		CreatedBy: updated_user.CreatedBy,
		UpdatedBy: updated_user.UpdatedBy,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "id is not a valid integer",
		})
		return
	}
	_, err = h.userUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "user not found",
		})
		return
	}
	_, err = h.userUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
