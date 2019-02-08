package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func GetTimestamp(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/timestamp", GetTimestamp).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}