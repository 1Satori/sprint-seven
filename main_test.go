package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body)

}

func TestMainHandlerWhenCNT(t *testing.T) {
	totalcnt := 4
	var bytes []byte

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cnt, _ := responseRecorder.Body.Read(bytes)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, len(strings.Split(string(cnt), ",")), totalcnt)

}

func TestMainHandlerCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=abs", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	require.NotEmpty(t, responseRecorder.Body)
}
