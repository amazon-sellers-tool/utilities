// Package snapshot Amazon snapshot response
package snapshot

// GetSnapshotResponse The Amazon Snapshot response
type GetSnapshotResponse struct {
	SnapshotID    string `json:"snapshotId"`
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
	Location      string `json:"location"`
	FileSize      int    `json:"fileSize"`
}
