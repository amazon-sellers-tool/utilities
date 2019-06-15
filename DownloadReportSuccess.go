// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// DownloadReportSuccess A success response from the API
func DownloadReportSuccess(w http.ResponseWriter, version int, result io.Reader) error {
	apiResponse := GetSnapshotAPIResponse{
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
