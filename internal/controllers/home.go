package controllers

import (
	"encoding/json"
	"net/http"
)

type data struct {
	Title string `json:"title`
}

type response struct {
	Code   int16 `json:"code"`
	Status bool  `json:"status"`
	Data   data  `json:"data"`
}

func Home(w http.ResponseWriter, r *http.Request) {

	resp := &response{
		Code: http.StatusOK,
		Data: data{
			Title: "Super secret area",
		},
		Status: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
