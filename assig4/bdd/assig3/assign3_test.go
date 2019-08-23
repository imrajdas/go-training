package main

import (
	"fmt"
	"maya/assig3/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateUsers(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/create", controller.Create).Methods("GET")
	req, err := http.NewRequest("GET", "/create", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestUpdateUsers(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/update", controller.Update).Methods("GET")
	req, err := http.NewRequest("GET", "/update", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/get/all", controller.GETALL).Methods("GET")
	req, err := http.NewRequest("GET", "/get/all", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestGetByIDUsers(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/get/1", controller.GETID).Methods("GET")
	req, err := http.NewRequest("GET", "/get/1", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/delete/1", controller.Delete).Methods("GET")
	req, err := http.NewRequest("GET", "/delete/1", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}
