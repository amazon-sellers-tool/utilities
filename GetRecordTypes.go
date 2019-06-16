// Package utilities Amazon Seller Utilities
package utilities

import (
	"log"
	"strings"
)

// GetRecordTypes Very important function that retrieves a token based on a profile id.
// @param string profileID Profile id
// @param string userID User id
// @return AmazonToken
func GetRecordTypes(adType string) recordTypes map[string]string {
	validRecordTypes := map[string]string{
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
	if !validRecordTypes[adType] {
		log.Output(1, "could not find a valid recordType\n")
		return nil
	}
	return validRecordTypes[adType]
}
