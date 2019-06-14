// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// RequestSnapshotSuccess A success response from the API
func RequestSnapshotSuccess(w http.ResponseWriter, version int, result io.Reader) error {
	var decodedBody RequestSnapshotResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		return SnapshotError(w, version, 404, errDecode)
	}
	apiResponse := RequestSnapshotAPIResponse{
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
