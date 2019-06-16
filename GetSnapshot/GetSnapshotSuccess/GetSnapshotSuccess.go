// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetSnapshotSuccess A success response from the API
func GetSnapshotSuccess(w http.ResponseWriter, version int, result io.Reader) utilities.GetSnapshotResponse {
	var decodedBody utilities.GetSnapshotResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		GetSnapshotError(w, version, 404, errDecode)
	}
	return decodedBody
}
