package data

import (
	"encoding/json"
	"net/http"
)

func getAllDates(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("my first go api")
}

func addNewDate(w http.ResponseWriter, r *http.Request) {

}

func removeDate(w http.ResponseWriter, r *http.Request) {

}