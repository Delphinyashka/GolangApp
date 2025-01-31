package services

import (
	"fmt"
	"net/http"
	"strings"
)

const authServiceURL = "http://localhost:8081"

func VerifyJWTToken(token string) (bool, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return false, fmt.Errorf("invalid token format")
	}

	req, err := http.NewRequest("GET", authServiceURL+"/user/verify", nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request to auth service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, fmt.Errorf("invalid or expired token, status code: %d", resp.StatusCode)
}
