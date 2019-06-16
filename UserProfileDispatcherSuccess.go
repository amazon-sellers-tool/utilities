// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// UserProfileDispatcherSuccess A success response from the API
func UserProfileDispatcherSuccess(w http.ResponseWriter, version int, results []string) error {
	apiResponse := UserProfileDispatcherAPIResponse{
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
