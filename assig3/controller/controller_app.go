package controller

import (
	"encoding/json"
	"fmt"
	"maya/assig3/domain"
	"maya/assig3/mapstore"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	store domain.CustomerStore
}

var controller = CustomerController{
	store: mapstore.NewMapStore(),
}

func Create(w http.ResponseWriter, r *http.Request) {

	customer := domain.Customer{
		ID:   "1",
		Name: "Raj Das",
	}

	err := controller.store.Create(customer)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(customer)

	w.Write(result)

}

func Update(w http.ResponseWriter, r *http.Request) {
	updateCustomer := domain.Customer{
		ID:   "1000",
		Name: "Deep Das",
	}
	err := controller.store.Update("1", updateCustomer)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("updated:", updateCustomer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(updateCustomer)

	w.Write(result)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	k := vars["id"]
	result := controller.store.Delete(k)
	if result != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(result)
	fmt.Fprintf(w, "Deleted")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rs, err := json.Marshal(result)

	w.Write(rs)

}

func GETID(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to Go Web Programming")
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	result, err := controller.store.GetByID(k)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rs, err := json.Marshal(result)

	w.Write(rs)
}

func GETALL(w http.ResponseWriter, r *http.Request) {

	result, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("GetAll:", result)
	fmt.Fprintf(w, "GetAll", result)
	rs, err := json.Marshal(result)

	w.Write(rs)

}
