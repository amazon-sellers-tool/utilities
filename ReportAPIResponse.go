// Package utilities Internal API response
package utilities

// ReportAPIResponse The Amazon Seller Tool response
type ReportAPIResponse struct {
	Version int            `json:"version"`
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Results ReportResponse `json:"results"`
	Error   string         `json:"error"`
}
