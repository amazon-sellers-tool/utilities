// Package utilities Amazon snapshot response
package utilities

// GetReportResponse The Amazon Report response
type GetReportResponse struct {
	ReportID      string `json:"reportId"`
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
	Location      string `json:"location"`
	FileSize      int    `json:"fileSize"`
}
