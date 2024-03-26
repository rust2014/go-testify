package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Корректный запрос
func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)                                   // статус 200
	require.NotEmpty(t, responseRecorder.Body.String(), "тело ответа не должно быть пустым") // проверка что тело ответа не пустое
}

// Неподдерживаемый город
func TestMainHandlerUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=spb", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code) // 400
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

// Запрос с count больше доступного
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	responseBody := responseRecorder.Body.String()

	list := strings.Split(responseBody, ",") // делим на элементы чтобы не счиать буквы
	require.Len(t, list, totalCount, "доступные кафе")
}
