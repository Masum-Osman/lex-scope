package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Masum-Osman/lex-scope/modules/text/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(service usecase.TextService) *gin.Engine {
	handler := NewTextHandler(service)
	r := gin.Default()
	api := r.Group("/api/v1")

	handler.RegisterRoutes(api)
	return r
}

func TestCreateText(t *testing.T) {
	service := usecase.NewTextService()
	router := setupRouter(service)

	payload := map[string]string{
		"content": "The quick brown fox jumps over the lazy dog.",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/texts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.EqualValues(t, 9, int(response["word_count"].(float64)))
}

func TestGetWordCount(t *testing.T) {
	service := usecase.NewTextService()
	router := setupRouter(service)

	req, _ := http.NewRequest("GET", "/api/v1/texts/dummy-id/word-count", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
