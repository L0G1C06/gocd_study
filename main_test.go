package main 

import (
	"net/http/httptest"
	"testing"
	"net/http"
)

func TestHelloServer(t *testing.T){
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloServer)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
	}

	expected := "Hello, World!"
	if rr.Body.String() != expected{
		t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)		
	}
}