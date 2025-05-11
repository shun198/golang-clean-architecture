package usecases_test

import (
	"testing"

	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	usecase "github.com/shun198/golang-clean-architecture/internal/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) GetOne(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) GetAll(params requests.ListUsersQuery) (*models.ListUsersResult, error) {
	args := m.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ListUsersResult), args.Error(1)
}

func (m *mockUserRepository) Update(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *mockUserRepository) DeleteOne(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

type UserUsecaseSuite struct {
	suite.Suite
	userUsecase *usecase.UserUsecase
	mockRepo    *mockUserRepository
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = NewMockUserRepository()
	suite.userUsecase = usecase.NewUserUsecase(suite.mockRepo)
}

func (suite *UserUsecaseSuite) TestGetUserSuccess() {
	// テスト用のユーザーID
	userID := 99

	// 期待値の設定
	expectedUser := &models.User{
		ID:       userID,
		Email:    "test_99@example.com",
		Username: "testuser99",
		Password: "hashedpassword",
		Role:     "admin",
	}

	// モックの設定
	suite.mockRepo.On("GetOne", userID).Return(expectedUser, nil)

	// テスト対象のメソッド実行
	user, err := suite.userUsecase.GetUser(userID)

	// 検証
	suite.Nil(err)
	suite.Equal(expectedUser, user)
	suite.mockRepo.AssertExpectations(suite.T())
}
