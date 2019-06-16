// Package report Amazon Seller Utilities API Responses
package report

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetReportSuccess A success response from the API
func GetReportSuccess(w http.ResponseWriter, version int, result io.Reader) GetReportResponse {
	var decodedBody GetReportResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		GetReportError(w, version, 404, errDecode)
	}
	return decodedBody
}
