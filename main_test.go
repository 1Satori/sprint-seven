package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status == http.StatusOK && assert.NotEmpty(t, responseRecorder.Body) {
	} else {
		t.Errorf("status not ok")
	}

	// здесь нужно добавить необходимые проверки
}

func TestMainHandlerWhenCNT(t *testing.T) {
	totalcnt := "4"

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	countStr := req.URL.Query().Get("count")

	if !assert.Equal(t, countStr, totalcnt) {
		httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	}
}

func TestMainHandlerCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cityStr := req.URL.Query().Get("city")

	if !assert.Equal(t, cityStr, "moscow") {
		t.Error("wrong city value")
		responseRecorder.Code = http.StatusBadRequest
	}
}
