package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request){

	user, err := getUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertOne(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func find(w http.ResponseWriter, r *http.Request){

	userId := mux.Vars(r)["id"]
	user := findOne(userId)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(user)

}

