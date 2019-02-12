package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	// "time"
	"github.com/gorilla/mux"

)

type Datestring struct {
	Date string `json:"date"`
	Unix int `json:"unix"`
	UTC string `json:"utc"`
}

// type Dates []Datestring

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/api", APIIndex).Methods("GET")
	router.HandleFunc("/api/", APIIndex).Methods("GET")
	router.HandleFunc("/api/timestamp", TimestampIndex).Methods("GET")
	router.HandleFunc("/api/timestamp/", TimestampShow).Methods("GET")
	router.HandleFunc("/api/timestamp/{datestring}", TimestampShow).Methods("GET")
	
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
