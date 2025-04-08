package middleware

import (
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		SuccessHandler: func(c echo.Context) {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			userID := claims["user_id"].(float64)
			slog.InfoContext(c.Request().Context(), "user_id", "user_id", userID)
			c.Set("user_id", int64(userID))
		},
	})
}
