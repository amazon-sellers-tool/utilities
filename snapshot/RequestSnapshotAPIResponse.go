// Package snapshot Internal API response
package snapshot

// RequestSnapshotAPIResponse The Amazon Seller Tool response
type RequestSnapshotAPIResponse struct {
	Version int                     `json:"version"`
	Success bool                    `json:"success"`
	Status  int                     `json:"status"`
	Results RequestSnapshotResponse `json:"results"`
	Error   string                  `json:"error"`
}
