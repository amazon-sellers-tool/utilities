// Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// SnapshotSuccess A success response from the API
func SnapshotSuccess(w http.ResponseWriter, version int, result io.Reader) {
	var decodedBody SnapshotResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		SnapshotError(w, version, 404, errDecode)
		return
	}
	apiResponse := ApiResponse{
		Version: version,
		Success: true,
		Status:  200,
		Results: decodedBody,
		Error:   nil,
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
