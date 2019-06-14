// Package utilities Amazon snapshot response
package utilities

// RequestReportResponse The Amazon Report response
type RequestReportResponse struct {
	ReportID      string `json:"reportId"`
	RecordType    string `json:"recordType"`
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
}
