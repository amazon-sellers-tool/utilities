// Package utilities Amazon Seller Utilities API Responses
package utilities

import (
	"encoding/json"
	"fmt"
)

// SnapshotSuccess A success response from the API
func SnapshotSuccess(body []byte) (*SnapshotResponse, error) {
	var s = new(SnapshotResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// SnapshotSuccess A success response from the API
// func SnapshotSuccess(w http.ResponseWriter, version int, body []byte) error {
// 	var decodedBody SnapshotResponse
// 	errDecode := json.NewDecoder(result).Decode(&decodedBody)
// 	if errDecode != nil {
// 		return SnapshotError(w, version, 404, errDecode)
// 	}
// 	apiResponse := SnapshotAPIResponse{
// 		Version: version,
// 		Success: true,
// 		Status:  200,
// 		Results: decodedBody,
// 		Error:   "",
// 	}
// 	apiResponseJSON, err := json.Marshal(apiResponse)
// 	if err != nil {
// 		log.Panic(err)
// 		panic(err)
// 	}
// 	log.Output(1, string(apiResponseJSON))
// 	return json.NewEncoder(w).Encode(apiResponse)
// }
