// Package report Internal API response
package report

// DownloadReportAPIResponse The Amazon Seller Tool response
type DownloadReportAPIResponse struct {
	Version int  `json:"version"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
	// Results DownloadReportResponse `json:"results"`
	Error string `json:"error"`
}
