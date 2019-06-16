// Package snapshot Amazon Seller Utilities API Responses
package snapshot

import (
	"encoding/json"
	"log"
	"net/http"
)

// DownloadSnapshotSuccess A success response from the API
func DownloadSnapshotSuccess(w http.ResponseWriter, version int) error {
	apiResponse := DownloadSnapshotAPIResponse{
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
