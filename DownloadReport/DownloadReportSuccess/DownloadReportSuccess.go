// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amazon-sellers-tool/utilities/DownloadReport/DownloadReportAPIResponse"
)

// DownloadReportSuccess A success response from the API
func DownloadReportSuccess(w http.ResponseWriter, version int) error {
	apiResponse := utilities.DownloadReportAPIResponse{
		Version: version,
		Success: true,
		Status:  200,
		// Results: decodedBody,
		Error: "",
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	return json.NewEncoder(w).Encode(apiResponse)
}
