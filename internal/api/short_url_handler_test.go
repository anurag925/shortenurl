package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/anurag/shortenurl/internal/api"
	"github.com/anurag/shortenurl/internal/db/models"
	"github.com/anurag/shortenurl/internal/dto"
	"github.com/anurag/shortenurl/internal/service"
	"github.com/anurag/shortenurl/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
	service.ShortURLService
}

func (m *mockService) ShortenURL(ctx context.Context, longURL string, customAlias *string, expiresAt *time.Time) (*models.ShortURL, error) {
	args := m.Called(ctx, longURL, customAlias, expiresAt)
	return args.Get(0).(*models.ShortURL), args.Error(1)
}

func (m *mockService) GetOriginalURL(ctx context.Context, shortCode string) (*models.ShortURL, error) {
	args := m.Called(ctx, shortCode)
	return args.Get(0).(*models.ShortURL), args.Error(1)
}

func TestShortURLHandler(t *testing.T) {
	e := echo.New()
	e.Validator = utils.NewCustomValidator()
	mockService := new(mockService)
	handler := api.NewShortURLHandler(mockService)

	t.Run("ShortenURL", func(t *testing.T) {
		longURL := "https://example.com"
		shortCode := "abc123"
		reqBody := dto.ShortenRequest{LongURL: longURL}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("ShortenURL", mock.Anything, longURL, (*string)(nil), (*time.Time)(nil)).
			Return(&models.ShortURL{
				ShortCode: shortCode,
				LongURL:   longURL,
			}, nil)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/shorten", bytes.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.ShortenURL(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response dto.ShortenResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.Equal(t, shortCode, response.ShortCode)
		assert.Equal(t, longURL, response.LongURL)
	})

	t.Run("Redirect", func(t *testing.T) {
		shortCode := "abc123"
		longURL := "https://example.com"

		mockService.On("GetOriginalURL", mock.Anything, shortCode).
			Return(&models.ShortURL{
				ShortCode: shortCode,
				LongURL:   longURL,
			}, nil)

		req := httptest.NewRequest(http.MethodGet, "/r/"+shortCode, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("short_code")
		c.SetParamValues(shortCode)

		err := handler.Redirect(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusMovedPermanently, rec.Code)
		assert.Equal(t, longURL, rec.Header().Get("Location"))
	})
}
