// Package report Internal API response
package report

// GetReportAPIResponse The Amazon Seller Tool response
type GetReportAPIResponse struct {
	Version int               `json:"version"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Results GetReportResponse `json:"results"`
	Error   string            `json:"error"`
}
