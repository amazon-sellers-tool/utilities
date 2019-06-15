// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GetReportSuccess A success response from the API
func GetReportSuccess(w http.ResponseWriter, version int, result io.Reader) error {
	var decodedBody GetReportResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		return GetReportError(w, version, 404, errDecode)
	}
	log.Output(1, string(apiResponseJSON))
	return decodedBody
}
