// Package utilities Amazon snapshot response
package utilities

// GetSnapshotResponse The Amazon Snapshot response
type GetSnapshotResponse struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		SnapshotID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"snapshotId"`
		RecordType struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"recordType"`
		Status struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"status"`
		StatusDetails struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"statusDetails"`
		Location struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"location"`
		FileSize struct {
			Description string `json:"description"`
			Type        int    `json:"type"`
		} `json:"fileSize"`
		Expiration struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"expiration"`
	} `json:"properties"`
	Code      string `json:"code"`
	Details   string `json:"details"`
	RequestID string `json:"requestId"`
}
