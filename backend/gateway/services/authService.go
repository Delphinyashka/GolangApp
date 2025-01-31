package services

import (
	"fmt"
	"net/http"
	"strings"
)

const authServiceURL = "http://localhost:8081" // Authentication service URL

// VerifyJWTToken sends the token to the authentication service for validation via the Authorization header
func VerifyJWTToken(token string) (bool, error) {
	// Ensure token starts with "Bearer "
	if !strings.HasPrefix(token, "Bearer ") {
		return false, fmt.Errorf("invalid token format")
	}

	// Create the request to the authentication service
	req, err := http.NewRequest("POST", authServiceURL+"/user/verify", nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %v", err)
	}

	// Add the Authorization header with the Bearer token
	req.Header.Add("Authorization", token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request to auth service: %v", err)
	}
	defer resp.Body.Close()

	// If status code is 200, the token is valid
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	// Return an error if the token is invalid or expired
	return false, fmt.Errorf("invalid or expired token, status code: %d", resp.StatusCode)
}
