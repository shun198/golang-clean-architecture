package usecase

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	repository "github.com/shun198/golang-clean-architecture/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	CreateUser(req requests.CreateUserRequest, auth_user_id int) (*models.User, error)
	GetUser(id int) (*models.User, error)
	GetAllUsers(params requests.ListUsersQuery) (*models.ListUsersResult, error)
	UpdateUser(req requests.UpdateUserRequest, user *models.User, auth_user_id int) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) GetAllUsers(params requests.ListUsersQuery) (*models.ListUsersResult, error) {
	return u.userRepository.GetAll(params)
}

func (u *UserUsecase) CreateUser(req requests.CreateUserRequest, auth_user_id int) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		Role:      req.Role,
		CreatedBy: auth_user_id,
		UpdatedBy: auth_user_id,
	}
	return u.userRepository.Create(user)
}

func (u *UserUsecase) GetUser(id int) (*models.User, error) {
	return u.userRepository.GetOne(id)
}

func (u *UserUsecase) UpdateUser(req requests.UpdateUserRequest, user *models.User, auth_user_id int) (*models.User, error) {
	user.Email = req.Email
	user.Username = req.Username
	user.Role = req.Role
	user.UpdatedBy = auth_user_id
	return u.userRepository.Update(user)
}

func (u *UserUsecase) DeleteUser(id int) (*models.User, error) {
	return u.userRepository.DeleteOne(id)
}
