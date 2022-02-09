package main

import (
	"lightning-talks/httptest/age"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(age.AgeHandler)
	log.Fatal(http.ListenAndServe(":1000", handler))
}
