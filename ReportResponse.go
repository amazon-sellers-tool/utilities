// Package utilities Amazon snapshot response
package utilities

// ReportResponse The Amazon Snapshot response
type ReportResponse struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		ReportID struct {
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"reportId"`
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
	} `json:"properties"`
	Code      string `json:"code"`
	Details   string `json:"details"`
	RequestID string `json:"requestId"`
}
