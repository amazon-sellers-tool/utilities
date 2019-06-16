// Package utilities Amazon Seller Utilities API Responses
package report

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// RequestReportSuccess A success response from the API
func RequestReportSuccess(w http.ResponseWriter, version int, result io.Reader) error {
	var decodedBody RequestReportResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		return RequestReportError(w, version, 404, errDecode)
	}
	apiResponse := RequestReportAPIResponse{
		Version: version,
		Success: true,
		Status:  200,
		Results: decodedBody,
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
