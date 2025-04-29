package usecase

import (
	"context"

	"github.com/shun198/golang-clean-architecture/internal/infrastructures/repositories"
)

type LoginUseCase interface {
	Execute(ctx context.Context, email, password string) (*TokenPair, error)
}

type loginUseCase struct {
	loginRepository repository.LoginRepository
}

func NewLoginUseCase(loginRepository repository.LoginRepository) LoginUseCase {
	return &loginUseCase{
		loginRepository: loginRepository,
	}
}

type JWTTokens struct {
	AccessToken  string
	RefreshToken string
}

var (
	ErrInvalidCredentials = repository.ErrInvalidCredentials
	ErrSystemError        = repository.ErrSystemError
)

func (u *loginUseCase) authenticateUser(ctx context.Context, email, password string) (*models.Operator, error) {
	user, err := u.loginRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// パスワードのハッシュを比較して認証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("認証に失敗しました")
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (u *loginUseCase) generateJWTTokens(user *models.User) (*JWTTokens, error) {
	now := time.Now()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  now.Add(consts.AccessTokenDuration).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": now.Add(consts.RefreshTokenDuration).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Printf("アクセストークンの生成に失敗しました: %v", err)
		return nil, ErrSystemError
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET_KEY")))
	if err != nil {
		log.Printf("リフレッシュトークンの生成に失敗しました: %v", err)
		return nil, ErrSystemError
	}

	return &JWTTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
