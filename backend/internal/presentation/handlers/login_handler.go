package handlers

import (
	"net/http"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shun198/golang-clean-architecture/internal/presentation/consts"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	usecase "github.com/shun198/golang-clean-architecture/internal/usecases"
)

var (
	ErrInvalidRequest = errors.New("リクエストの形式が正しくありません")
)

// クッキーを設定する共通関数
func setAuthCookie(c *gin.Context, name, value string, maxAge time.Duration) {
	config := consts.NewCookieConfig()

	httpOnly := config.HttpOnly
	if name == "access_token" {
		// クライアントサイドでのAPIリクエスト認証処理のために必要（最善ではないが、検証なので今回はこれで妥協。実際はapi routesなどを経由させる）
		httpOnly = false
	}

	c.SetCookie(
		name,                  // Name: クッキー名
		value,                 // Value: クッキーの値
		int(maxAge.Seconds()), // MaxAge: 有効期限（秒）
		consts.CookiePath,     // Path: クッキーの有効範囲
		config.Domain,         // Domain: クッキーが有効なドメイン
		config.Secure,         // Secure: HTTPSのみ
		httpOnly,              // HttpOnly: JSからのアクセス制限
	)
	c.SetSameSite(config.SameSite) // SameSite属性を別途設定
}

// LoginHandlerの構造体を追加
type LoginHandler struct {
	loginUseCase usecase.LoginUseCase
}

// NewLoginHandlerコンストラクタを追加
func NewLoginHandler(loginUseCase usecase.LoginUseCase) *LoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}
}

// ユーザーのログイン処理を行う
func (h *LoginHandler) Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "validation error",
		})
		return
	}

	tokenPair, err := h.loginUseCase.Execute(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed creating token",
		})
		return
	}

	h.setAuthCookies(c, tokenPair)
	c.JSON(http.StatusOK, gin.H{
		"msg": "login success",
	})
}

// 認証トークンをクッキーに設定する
func (h *LoginHandler) setAuthCookies(c *gin.Context, tokenPair *usecase.TokenPair) {
	setAuthCookie(c, consts.AccessTokenCookie, tokenPair.AccessToken, consts.AccessTokenDuration)
	setAuthCookie(c, consts.RefreshTokenCookie, tokenPair.RefreshToken, consts.RefreshTokenDuration)
}
