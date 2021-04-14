package main

import (
	"encoding/json"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	body := map[string]string{"message":"OK"}
	json.NewEncoder(w).Encode(body)
}

func getUserFromRequest(r *http.Request) (User, error) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	return user, err
}

