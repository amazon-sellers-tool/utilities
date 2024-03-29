// Package snapshot Amazon Seller Utilities API Responses
package snapshot

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestSnapshotError An error response from the API
func RequestSnapshotError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := RequestSnapshotAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: RequestSnapshotResponse{},
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
