package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Success bool    `json:"Success"`
	ErrCode string  `json:"ErrCode"`
	Value   float64 `json:"Value"`
}

func main() {
	router := mux.NewRouter()
	// regexp: all digits with one symbol after .
	router.HandleFunc("/api/{method}/{a:[+-]?[0-9]*[.]?[0-9]}/{b:[+-]?[0-9]*[.]?[0-9]}", apiService).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func convertToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func apiService(w http.ResponseWriter, r *http.Request) {
	var mathResult float64

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	a, b := convertToFloat(params["a"]), convertToFloat(params["b"])

	switch params["method"] {
	case "add":
		mathResult = a + b
	case "sub":
		mathResult = a - b
	case "mul":
		mathResult = a * b
	case "div":
		if b == 0 {
			json.NewEncoder(w).Encode(&Response{Success: false, ErrCode: "Division by zero."})
			return
		}
		mathResult = a / b
	default:
		json.NewEncoder(w).Encode(&Response{Success: false, ErrCode: "Unknown method."})
		return
	}

	json.NewEncoder(w).Encode(&Response{
		Success: true,
		ErrCode: "",
		Value:   mathResult,
	})
	return
}
