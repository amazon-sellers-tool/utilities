// Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"net/http"
)

type SnapshotResponse struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		SnapshotId struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"snapshotId"`
		RecordType struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"recordType"`
		Status struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"status"`
		StatusDetail struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"statusDetails"`
		Location struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"location"`
		FileSize struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"fileSize"`
		Expiration struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"expiration"`
	} `json:"properties"`
	Code      string `json:"code"`
	Details   string `json:"details"`
	RequestId string `json:"requestId"`
}

type ApiResponse struct {
	Version int              `json:"version"`
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Results SnapshotResponse `json:"results"`
	Error   error            `json:"error"`
}

func SnapshotError(w http.ResponseWriter, version int, code int, err error) {
	apiResponse := ApiResponse{
		Version: version,
		Success: true,
		Status:  code,
		Results: SnapshotResponse{},
		Error:   err,
	}
	json.NewEncoder(w).Encode(apiResponse)
	return
}
