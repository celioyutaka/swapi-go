package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

//default header like Content-Type as application/json
func SetHeaderAPI(w http.ResponseWriter) http.ResponseWriter {

	w.Header().Set("Content-Type", "application/json")
	return w
}
func setResponse(w http.ResponseWriter, data interface{}, result bool) {
	response := Response{Data: data, Success: result}
	json.NewEncoder(w).Encode(response)
}
