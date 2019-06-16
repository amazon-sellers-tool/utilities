// Package utilities Internal API response
package utilities

// RequestSnapshotAPIResponse The Amazon Seller Tool response
type RequestSnapshotAPIResponse struct {
	Version int               `json:"version"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Results map[string]string `json:"results"`
	Error   string            `json:"error"`
}
