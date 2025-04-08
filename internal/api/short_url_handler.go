package api

import (
	"log/slog"
	"net/http"

	"github.com/anurag/shortenurl/internal/dto"
	"github.com/anurag/shortenurl/internal/service"
	"github.com/labstack/echo/v4"
)

type ShortURLHandler struct {
	service service.ShortURLService
}

func NewShortURLHandler(service service.ShortURLService) *ShortURLHandler {
	return &ShortURLHandler{service: service}
}

// ShortenURL godoc
//
//	@Summary		Shorten a URL
//	@Description	Create a short URL from a long URL
//	@Tags			shortener
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.ShortenRequest	true	"URL to shorten"
//	@Success		201		{object}	dto.ShortenResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/v1/shorten [post]
func (h *ShortURLHandler) ShortenURL(c echo.Context) error {
	req := new(dto.ShortenRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	shortURL, err := h.service.ShortenURL(c.Request().Context(), req.LongURL, req.CustomAlias, req.ExpiresAt)
	if err != nil {
		slog.Error("Failed to shorten URL", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to shorten URL")
	}

	response := dto.ShortenResponse{
		ShortCode: shortURL.ShortCode,
		LongURL:   shortURL.LongURL,
		ShortURL:  c.Scheme() + "://" + c.Request().Host + "/r/" + shortURL.ShortCode,
		ExpiresAt: shortURL.ExpiresAt,
	}

	return c.JSON(http.StatusCreated, response)
}

// Redirect godoc
//
//	@Summary		Redirect to original URL
//	@Description	Redirects to the original URL from a short code
//	@Tags			shortener
//	@Param			short_code	path	string	true	"Short code"
//	@Success		301
//	@Failure		404	{object}	map[string]string
//	@Router			/r/{short_code} [get]
func (h *ShortURLHandler) Redirect(c echo.Context) error {
	shortCode := c.Param("short_code")
	if shortCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Short code is required")
	}

	shortURL, err := h.service.GetOriginalURL(c.Request().Context(), shortCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "URL not found")
	}

	return c.Redirect(http.StatusMovedPermanently, shortURL.LongURL)
}
