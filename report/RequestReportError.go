// Package utilities Amazon Seller Utilities API Responses
package report

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestReportError An error response from the API
func RequestReportError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := RequestReportAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: RequestReportResponse{},
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
