// Package utilities Internal API response
package utilities

// GetSnapshotAPIResponse The Amazon Seller Tool response
type GetSnapshotAPIResponse struct {
	Version int                 `json:"version"`
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Results GetSnapshotResponse `json:"results"`
	Error   string              `json:"error"`
}
