package handlers

import (
	"Number-Classification-API/middleware"
	"Number-Classification-API/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type APIResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

func ClassifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	middleware.HandleCORS()

	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, `{"error" : "number parameter is required"}`, http.StatusBadRequest)
		return
	}
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, `{"number": "alphabet", "error": true}`, http.StatusBadRequest)
		return
	}

	// More validation logic are added here
	response := APIResponse{
		Number:     number,
		IsPrime:    IsPrime(number),
		IsPerfect:  IsPerfect(number),
		Properties: []string{},
		DigitSum:   SumOfDigits(number),
		FunFact:    "",
	}

	if IsArmStrong(number) {
		response.Properties = append(response.Properties, "armstrong")
	}
	if IsEven(number) {
		response.Properties = append(response.Properties, "even")
	} else {
		response.Properties = append(response.Properties, "odd")
	}

	funFact, err := utils.FetchFunFact(number)
	if err != nil {
		http.Error(w, `{"error": "unable to fetch fun fact"}`, http.StatusInternalServerError)
		return
	}
	response.FunFact = funFact

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
