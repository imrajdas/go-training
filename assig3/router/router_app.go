package router

import (
	"maya/assig3/controller"

	"github.com/gorilla/mux"
)

var Mux = mux.NewRouter()

// type CustomerController struct {
// 	store domain.CustomerStore
// }

// var ccontroller = CustomerController{
// 	store: mapstore.NewMapStore(),
// }

func Routers() {

	Mux.HandleFunc("/create", controller.Create).Methods("GET")
	Mux.HandleFunc("/update", controller.Update).Methods("GET")
	Mux.HandleFunc("/get/{id}", controller.GETID).Methods("GET")
	Mux.HandleFunc("/get/all", controller.GETALL).Methods("GET")
	Mux.HandleFunc("/delete/{id}", controller.Delete).Methods("GET")

}
