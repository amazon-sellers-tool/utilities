// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetSnapshotError An error response from the API
func GetSnapshotError(w http.ResponseWriter, version int, code int, err error) error {
	apiResponse := GetSnapshotAPIResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: GetSnapshotResponse{},
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
