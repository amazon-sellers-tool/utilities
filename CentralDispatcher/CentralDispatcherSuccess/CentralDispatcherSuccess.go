// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/amazon-sellers-tool/utilities/CentralDispatcher/CentralDispatcherAPIResponse"
)

// CentralDispatcherSuccess A success response from the API
func CentralDispatcherSuccess(w http.ResponseWriter, version int, results map) error {
	apiResponse := utilities.CentralDispatcherAPIResponse{
		Version: version,
		Success: true,
		Status:  200,
		Results: results,
	}
	apiResponseJSON, err := json.Marshal(apiResponse)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Output(1, string(apiResponseJSON))
	return json.NewEncoder(w).Encode(apiResponse)
}
