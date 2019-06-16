// Package utilities Amazon Seller Utilities
package utilities

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// GetToken Very important function that retrieves a token based on a profile id.
// @param string profileID Profile id
// @param string userID User id
// @return AmazonToken
func TestGetToken(profileID string, userID string) (token string, err error) {
	payload := map[string]string{
		"profileId": profileID,
		"userId":    userID,
	}
	jsonValue, _ := json.Marshal(payload)
	log.Output(1, string(jsonValue))

	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("GET_AMAZON_TOKEN_SERVICE_ADDRESS"),
		bytes.NewBuffer(jsonValue),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", os.Getenv("GET_AMAZON_TOKEN_SERVICE_TOKEN"))

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var amazonToken AmazonToken

	err = json.NewDecoder(resp.Body).Decode(&amazonToken)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	req.Body.Close()
	return amazonToken.AccessToken, nil
}
