// Package report Internal API response
package report

// RequestReportAPIResponse The Amazon Seller Tool response
type RequestReportAPIResponse struct {
	Version int                   `json:"version"`
	Success bool                  `json:"success"`
	Status  int                   `json:"status"`
	Results RequestReportResponse `json:"results"`
	Error   string                `json:"error"`
}
