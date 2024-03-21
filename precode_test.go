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

	assert.Equal(t, http.StatusOK, responseRecorder.Code) // статус 200
	responseBody := responseRecorder.Body.String()

	expectedCafes := []string{"Мир кофе", "Сытый студент"}
	list := strings.Split(responseBody, ",") // делим на элементы чтобы не счиать буквы
	assert.ElementsMatch(t, expectedCafes, list)
}

// Неподдерживаемый город
func TestMainHandlerUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1&city=spb", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code) // 400
	assert.Contains(t, responseRecorder.Body.String(), "неправильное значение города")
}

// Запрос с count больше доступного
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	responseBody := responseRecorder.Body.String()

	require.NotEmpty(t, responseBody)
	list := strings.Split(responseBody, ",") // делим на элементы чтобы не счиать буквы
	assert.Len(t, list, totalCount, "доступные кафе")

}
