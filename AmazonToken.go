// Package utilities Internal Amazon Token struct
package utilities

import "time"

// AmazonToken Describes what the token will look like
type AmazonToken struct {
	AccessToken       string        `json:"amazon_access_token"`
	AccessTokenExpiry time.Duration `json:"amazon_accessToken_expiry"`
}
