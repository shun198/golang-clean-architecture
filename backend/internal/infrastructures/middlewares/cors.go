package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() gin.HandlerFunc {
	allowOrigins := strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ",")
	allowMethods := strings.Split(os.Getenv("CORS_ALLOW_METHODS"), ",")
	allowHeaders := strings.Split(os.Getenv("CORS_ALLOW_HEADERS"), ",")
	exposeHeaders := strings.Split(os.Getenv("CORS_EXPOSE_HEADERS"), ",")
	allowCredentials := os.Getenv("CORS_ALLOW_CREDENTIALS") == "true"
	maxAge, err := time.ParseDuration(os.Getenv("CORS_MAX_AGE"))
	if err != nil {
		maxAge = 12 * time.Hour
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     allowMethods,
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
	})
}
