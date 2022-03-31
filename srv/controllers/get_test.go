package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSayHi(t *testing.T) {
	path := "/v1/say/hi"
	router := gin.Default()
	ctrl := new(Controllers)

	router.GET(path, ctrl.SayHi)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rec.Body.String() != `{"err":null,"message":"Hi"}` {
		t.Errorf("handler returned wrong status code: got %v want %v", rec.Body.String(), `{"err":null,"message":"Hi"}`)
	}
}
