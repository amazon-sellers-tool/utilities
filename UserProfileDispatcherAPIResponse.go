// Package utilities Internal API response
package utilities

// UserProfileDispatcherAPIResponse The Amazon Seller Tool response
type UserProfileDispatcherAPIResponse struct {
	Version int      `json:"version"`
	Success bool     `json:"success"`
	Status  int      `json:"status"`
	Results []string `json:"results"`
	Error   string   `json:"error"`
}
