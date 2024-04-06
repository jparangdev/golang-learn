package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	ast.Nil(err)
	ast.Equal(2, len(list))
	ast.Equal("jude", list[0].Name)
	ast.Equal("heesoo", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	var student Student
	ast.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	ast.Nil(err)
	ast.Equal("jude", student.Name)
	ast.Equal(19, student.Age)
}

func TestJsonHandler3(t *testing.T) {
	ast := assert.New(t)

	var student Student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(`{"id":0,"Name":"haejoo","Age":15,"Score":99}`))

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	ast.Nil(err)
	ast.Equal("haejoo", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	ast := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)

	ast.Equal(http.StatusNotFound, res.Code)
}
