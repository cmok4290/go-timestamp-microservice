package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "time"
	"github.com/gorilla/mux"

)

type Datestring struct {
	Date string `json:"date"`
	Unix int `json:"unix"`
	UTC string `json:"utc"`
}

// type Dates []Datestring

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
	w.WriteHeader(http.StatusOK)
}

func APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add API documentation...\n"))
	w.WriteHeader(http.StatusOK)
}

func TimestampIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Another index..."))
	w.WriteHeader(http.StatusOK)
}

func TimestampShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ds := vars["datestring"]

	date := Datestring{Date: ds, Unix: 5, UTC: "UTC"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(date); err != nil {
		panic(err)
	}

	// fmt.Fprintln(w, "Timestamp shows:", ds)
}

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
