// Package utilities Amazon snapshot response
package utilities

// RequestSnapshotResponse The Amazon Snapshot response
type RequestSnapshotResponse struct {
	SnapshotID    string `json:"snapshotId"`
	RecordType    string `json:"recordType"`
	Status        string `json:"status"`
	StatusDetails string `json:"statusDetails"`
}
