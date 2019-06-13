// Package utitlities
package utilities

// SnapshotResponse The Amazon Snapshot response
type SnapshotResponse struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		SnapshotId struct {
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
		StatusDetail struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"statusDetails"`
		Location struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"location"`
		FileSize struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"fileSize"`
		Expiration struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"expiration"`
	} `json:"properties"`
	Code      string `json:"code"`
	Details   string `json:"details"`
	RequestId string `json:"requestId"`
}
