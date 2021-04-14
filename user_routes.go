package main

import (
	"github.com/gorilla/mux"
)

func getRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ping).Methods("GET")
	router.HandleFunc("/users", create).Methods("POST")
	router.HandleFunc("/users/{id}", find).Methods("GET")
	return router
}

