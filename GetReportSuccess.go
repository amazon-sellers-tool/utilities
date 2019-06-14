// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetReportSuccess A success response from the API
func GetReportSuccess(w http.ResponseWriter, version int, result io.Reader) error {
	var decodedBody GetReportResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		return GetSnapshotError(w, version, 404, errDecode)
	}
	apiResponse := GetReportAPIResponse{
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
