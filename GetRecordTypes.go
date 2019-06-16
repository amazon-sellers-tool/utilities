// Package utilities Amazon Seller Utilities
package utilities

import (
	"log"
)

// GetRecordTypes Very important function that retrieves a token based on a profile id.
// @param string profileID Profile id
// @param string userID User id
// @return AmazonToken
func GetRecordTypes(adType string) []string {
	validRecordTypes := map[string][]string{
		"sp": []string{
			"campaigns",
			"adGroups",
			"keywords",
			"productAds",
			"targets",
		},
		"hsa": []string{
			"campaigns",
			"adGroups",
			"keywords",
		},
	}

	// Check the validity of record type
	if val, ok := validRecordTypes[adType]; ok {
		return val
	}

	log.Output(1, "could not find a valid recordType\n")
	return nil
}
