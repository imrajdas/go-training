package main

import (
	"log"
	"maya/assig3/router"
	"net/http"
)

func main() {
	log.Println("Listening...")
	router.Routers()
	http.ListenAndServe(":8080", router.Mux)
}
