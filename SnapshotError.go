// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// SnapshotError An error response from the API
func SnapshotError(w http.ResponseWriter, version int, code int, err error) {
	apiResponse := ApiResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: SnapshotResponse{},
		Error:   err,
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	json.NewEncoder(w).Encode(apiResponse)
	return
}
