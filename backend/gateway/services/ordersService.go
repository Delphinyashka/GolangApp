package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const ordersServiceURL = "http://localhost:8084"

func FetchOrders(orderIDs []string, page int, limit int) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/orders?ids=%s&page=%d&limit=%d", ordersServiceURL, strings.Join(orderIDs, ","), page, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders from orders service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch orders: status code %d", resp.StatusCode)
	}

	var orders []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return orders, nil
}

func CalculateTotalAmount() (float64, error) {
	url := fmt.Sprintf("%s/total", ordersServiceURL)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to get total amount from orders service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch total amount: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if totalAmount, ok := result["totalAmount"].(float64); ok {
		return totalAmount, nil
	}
	return 0, fmt.Errorf("totalAmount not found in the response")
}
