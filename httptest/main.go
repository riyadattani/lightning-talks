package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(AgeServer)
	log.Fatal(http.ListenAndServe(":1000", handler))
}

func AgeServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "21")
}
