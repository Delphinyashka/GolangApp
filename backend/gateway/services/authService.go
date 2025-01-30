package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const authServiceURL = "http://localhost:8081" // Authentication service URL

// VerifyJWTToken sends the token to the authentication service for validation
func VerifyJWTToken(token string) (bool, error) {
	// Create the request to the authentication service
	requestBody, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		return false, fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := http.Post(authServiceURL+"/user/verify", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, fmt.Errorf("failed to send request to auth service: %v", err)
	}
	defer resp.Body.Close()

	// If status code is 200, the token is valid
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
