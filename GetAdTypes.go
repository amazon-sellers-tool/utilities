// Package utilities Amazon Seller Utilities
package utilities

// GetAdTypes Very important function that retrieves a token based on a profile id.
// @param string profileID Profile id
// @param string userID User id
// @return AmazonToken
func GetAdTypes() []string {
	adTypes := []string{
		"sp",
		"hsa",
	}
	return adTypes
}
