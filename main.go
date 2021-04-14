package main

import (
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer()  {
	router := getRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}

