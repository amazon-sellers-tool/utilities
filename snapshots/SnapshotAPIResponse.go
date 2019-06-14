// Package utilities Internal API response
package utilities

// SnapshotAPIResponse The Amazon Seller Tool response
type SnapshotAPIResponse struct {
	Version int              `json:"version"`
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Results SnapshotResponse `json:"results"`
	Error   string           `json:"error"`
}
