// Package utilities Internal API response
package utilities

// GetReportAPIResponse The Amazon Seller Tool response
type GetReportAPIResponse struct {
	Version int               `json:"version"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Results GetReportResponse `json:"results"`
	Error   string            `json:"error"`
}
