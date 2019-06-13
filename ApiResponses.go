// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ApiResponse The interface
type APIResponse interface {
	SnapshotError(w http.ResponseWriter, version int, code int, err error)
	SnapshotSuccess(w http.ResponseWriter, version int, result io.Reader)
}

// APIResponses holds funcs
type APIResponses struct{}

// SnapshotResponse See https://advertising.amazon.com/API/docs/v2/reference/snapshots
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

// Response An basic api response struct
type Response struct {
	Version int              `json:"version"`
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Results SnapshotResponse `json:"results"`
	Error   error            `json:"error"`
}

// SnapshotError An error response from the API
func (APIResponses) SnapshotError(w http.ResponseWriter, version int, code int, err error) {
	apiResponse := Response{
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

// SnapshotSuccess A success response from the API
func (APIResponses) SnapshotSuccess(w http.ResponseWriter, version int, result io.Reader) {
	var decodedBody SnapshotResponse
	errDecode := json.NewDecoder(result).Decode(&decodedBody)
	if errDecode != nil {
		GetAPIResponses().SnapshotError(w, version, 404, errDecode)
		return
	}
	apiResponse := Response{
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

// GetAPIResponses
func GetAPIResponses() APIResponse {
	return APIResponses{}
}
