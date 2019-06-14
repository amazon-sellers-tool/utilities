// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetReportError An error response from the API
func GetReportError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := GetReportAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: GetReportResponse{},
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
