package main

import (
	"bytes"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBingForm(t *testing.T) {
	router := setupRouter()
	//data := `{"username":"admin", "password": "123456"}`
	jsonStr := []byte(`{"username": {"admin"}, "password": {"123456"}}`)
	// mock 一个HTTP请求
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	// mock 一个response响应器(ResponseRecorde)
	resp := httptest.NewRecorder()
	// 让server端处理  req 并记录返回的响应内容到 resp
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

}
