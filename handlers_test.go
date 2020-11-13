package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestHandleRoot(t *testing.T) {
	is := is.New(t)
	handler := getRoutes()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, req)

	is.Equal(http.StatusOK, rr.Code)
	resBytes, err := ioutil.ReadAll(rr.Body)
	is.NoErr(err)
	is.Equal("<h2> Hello from a Go server </h2>", string(resBytes))
}

func TestHandleGreeting(t *testing.T) {
	is := is.New(t)
	handler := getRoutes()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/John", nil)
	handler.ServeHTTP(rr, req)

	is.Equal(http.StatusOK, rr.Code)
	resBytes, err := ioutil.ReadAll(rr.Body)
	is.NoErr(err)
	is.Equal("<p> Hello John </p>", string(resBytes))
}
