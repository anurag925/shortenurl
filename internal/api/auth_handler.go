package api

import (
	"net/http"

	"github.com/anurag/shortenurl/internal/dto"
	"github.com/anurag/shortenurl/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	req := new(dto.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	token, err := h.authService.Register(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.AuthResponse{Token: token})
}

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(dto.LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	token, err := h.authService.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}

// Add this to AuthHandler
func (h *AuthHandler) CheckAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"authenticated": true})
}
