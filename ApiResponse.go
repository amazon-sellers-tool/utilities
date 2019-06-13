// Package utilities Internal API response
package utilities

// APIResponse The Amazon Seller Tool response
type APIResponse struct {
	Version int              `json:"version"`
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Results SnapshotResponse `json:"results"`
	Error   error            `json:"error"`
}
