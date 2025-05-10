package usecase

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/consts"
	repository "github.com/shun198/golang-clean-architecture/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

// インターフェースを定義
type LoginUseCase interface {
	Execute(ctx context.Context, email, password string) (*TokenPair, error)
}

// インターフェースを実装する構造体
type loginUseCase struct {
	loginRepository repository.ILoginRepository
}

// 新しい loginUseCase を作成する関数
func NewLoginUseCase(loginRepository repository.ILoginRepository) LoginUseCase {
	return &loginUseCase{
		loginRepository: loginRepository,
	}
}

// アクセストークンとリフレッシュトークンを保持
type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// ユーザーの認証とトークンを生成
func (u *loginUseCase) Execute(ctx context.Context, email, password string) (*TokenPair, error) {
	user, err := u.authenticateUser(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return u.generateTokenPair(user)
}

// ユーザーの認証を行う
func (u *loginUseCase) authenticateUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := u.loginRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// パスワードのハッシュを比較して認証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("認証に失敗しました")
		return nil, err
	}

	return user, nil
}

// アクセストークンとリフレッシュトークンを生成
func (u *loginUseCase) generateTokenPair(user *models.User) (*TokenPair, error) {
	now := time.Now()

	// アクセストークンを生成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  now.Add(consts.AccessTokenDuration).Unix(),
	})

	// リフレッシュトークンを生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": now.Add(consts.RefreshTokenDuration).Unix(),
	})

	// アクセストークンを署名
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Printf("アクセストークンの生成に失敗しました: %v", err)
		return nil, err
	}

	// リフレッシュトークンを署名
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET_KEY")))
	if err != nil {
		log.Printf("リフレッシュトークンの生成に失敗しました: %v", err)
		return nil, err
	}

	// トークンペアを返却
	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
