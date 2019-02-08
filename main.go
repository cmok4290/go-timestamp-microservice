package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	// "time"
	"github.com/gorilla/mux"

)

type Datestring struct {
	date string `json:"datestring"`
	unix int `json:"unix"`
	utc string `json:"utc"`
}

type Dates []Datestring

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
}

func APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add API documentation...\n"))
}

func TimestampIndex(w http.ResponseWriter, r *http.Request) {
	return
}

func TimestampShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ds := vars["datestring"]
	fmt.Fprintln(w, "Timestamp shows:", ds)
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