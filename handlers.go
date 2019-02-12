package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"regexp"
	// "strconv"
	// "time"
	"github.com/gorilla/mux"
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
	w.Write([]byte("Another index...\n"))
	w.WriteHeader(http.StatusOK)
}

func TimestampShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ds := vars["datestring"]
	// ms := uint64(0)

	re := regexp.MustCompile("[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}")
	if re.MatchString(ds) != true {
		// if s, err := strconv.ParseUint(ds, 10, 64); err == nil {
		// 	ms = s
		// }
		// ms = s
	}

	date := Datestring{Date: ds, Unix: 5, UTC: "UTC"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(date); err != nil {
		panic(err)
	}

	// fmt.Fprintln(w, "Timestamp shows:", ds)
}
