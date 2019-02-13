package main

type Datestring struct {
	Date string `json:"date"`
	Unix int `json:"unix"`
	UTC string `json:"utc"`
	Natural string `json:"natural"`
}
