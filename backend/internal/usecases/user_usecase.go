package usecase

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	repository "github.com/shun198/golang-clean-architecture/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	CreateUser(requests.CreateUserRequest) (*models.User, error)
	GetUser(id int) (*models.User, error)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUser(req requests.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	return u.userRepository.Create(user)
}

func (u *UserUsecase) GetUser(id int) (*models.User, error) {
	return u.userRepository.GetOne(id)
}
