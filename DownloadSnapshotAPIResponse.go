// Package utilities Internal API response
package utilities

// DownloadSnapshotAPIResponse The Amazon Seller Tool response
type DownloadSnapshotAPIResponse struct {
	Version int  `json:"version"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
	// Results DownloadSnapshotResponse `json:"results"`
	Error string `json:"error"`
}
