// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amazon-sellers-tool/utilities/DownloadSnapshot/DownloadSnapshotAPIResponse"
)

// DownloadSnapshotError An error response from the API
func DownloadSnapshotError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := utilities.DownloadSnapshotAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		// Results: DownloadSnapshotResponse{},
		Error: err.Error(),
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	return json.NewEncoder(w).Encode(apiResponse)
}
