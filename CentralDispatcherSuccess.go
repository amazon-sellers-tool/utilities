// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// CentralDispatcherSuccess A success response from the API
func CentralDispatcherSuccess(w http.ResponseWriter, version int, results map[string]string) error {
	apiResponse := CentralDispatcherAPIResponse{
		Version: version,
		Success: true,
		Status:  200,
		Results: results,
		Error:   "",
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	return json.NewEncoder(w).Encode(apiResponse)
}
