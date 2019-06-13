// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// ReportError An error response from the API
func ReportError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := ReportAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: ReportResponse{},
		Error:   err.Error(),
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	return json.NewEncoder(w).Encode(apiResponse)
}
