// Package utilities Internal API response
package utilities

// RequestReportAPIResponse The Amazon Seller Tool response
type RequestReportAPIResponse struct {
	Version int                   `json:"version"`
	Success bool                  `json:"success"`
	Status  int                   `json:"status"`
	Results RequestReportResponse `json:"results"`
	Error   string                `json:"error"`
}
