package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/{method}/{a:[+-]?[0-9]*[.]?[0-9]}/{b:[+-]?[0-9]*[.]?[0-9]}", apiService).Methods("GET")
	return router
}

func TestApiServiceUnknownMethod(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/unknown/1/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":false,"ErrCode":"Unknown method.","Value":0}`)
}

func TestApiServiceNotDigit(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1/a", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "404 response is expected")
}

// more than one symbol after dot in float like 0.53
// only 0.5 allowed
func TestApiServiceAfterDot(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1/1.12", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "404 response is expected")
}

//-------------------------------- add
func TestApiServiceAdd(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":2}`)
}

func TestApiServiceAddMinusPlus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/-1/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":0}`)
}

func TestApiServiceAddPlusMinus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1/-1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":0}`)
}

func TestApiServiceAddFloatInteger(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1.2/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":2.2}`)
}

func TestApiServiceAddIntegerFloat(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/add/1/1.2", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":2.2}`)
}

//------------------------------------ sub
func TestApiServiceSub(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/sub/1/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":0}`)
}

func TestApiServiceSubMinusPlus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/sub/-1/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-2}`)
}

func TestApiServiceSubPlusMinus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/sub/1/-1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":2}`)
}

func TestApiServiceSubFloatInteger(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/sub/1.2/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":0.19999999999999996}`)
}

func TestApiServiceSubIntegerFloat(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/sub/1/1.2", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-0.19999999999999996}`)
}

//------------------------------------ mul
func TestApiServiceMul(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/mul/5/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":15}`)
}

func TestApiServiceMulMinusPlus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/mul/-5/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-15}`)
}

func TestApiServiceMulPlusMinus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/mul/5/-3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-15}`)
}

func TestApiServiceMulFloatInteger(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/mul/3.9/5", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":19.5}`)
}

func TestApiServiceMulIntegerFloat(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/mul/5/3.7", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":18.5}`)
}

//------------------------------------ div
func TestApiServiceDiv(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/5/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":1.6666666666666667}`)
}

func TestApiServiceDivMinusPlus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/-5/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-1.6666666666666667}`)
}

func TestApiServiceDivPlusMinus(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/5/-7", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":-0.7142857142857143}`)
}

func TestApiServiceDivFloatInteger(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/3.9/5", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":0.78}`)
}

func TestApiServiceDivIntegerFloat(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/5/3.7", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":true,"ErrCode":"","Value":1.3513513513513513}`)
}

func TestApiServiceDivZero(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/div/5/0", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
	assert.JSONEq(t, response.Body.String(), `{"Success":false,"ErrCode":"Division by zero.","Value":0}`)
}
