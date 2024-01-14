package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func getAllDates(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all dates")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("my first go api")
}

func handleRequests() {
	http.Handle("/getDates", http.HandlerFunc(getAllDates))
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	handleRequests()
}