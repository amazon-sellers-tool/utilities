// Package utitlities
package utilities

// ApiResponse The Amazon Seller Tool response
type ApiResponse struct {
	Version int              `json:"version"`
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Results SnapshotResponse `json:"results"`
	Error   error            `json:"error"`
}
