package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBarHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=test&id=1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)
	expected := "Hello test, your ID is 1"
	ast.Equal(expected, res.Body.String())
}

func TestJsonHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student)
	ast.Nil(err)
	expected := &Student{"jude", 35, 95}
	ast.Equal(expected, student)
}
