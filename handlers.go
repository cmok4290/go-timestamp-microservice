package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"regexp"
	// "strconv"
	"github.com/gorilla/mux"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
	w.WriteHeader(http.StatusOK)
}

func APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add API documentation...\n"))
	w.WriteHeader(http.StatusOK)
}

func TimestampIndex(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	date := Datestring{
		Unix: t.Unix(), UTC: t.UTC().String(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(date); err != nil {
		panic(err)
	}
}

func TimestampShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ds := vars["datestring"]

	// format: "Fri, 25 Dec 2015 00:00:00 GMT" aka RFC1123
	// error: "Invalid Date"
	date := Datestring{Unix: 5, UTC: "UTC"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(date); err != nil {
		panic(err)
	}
}
