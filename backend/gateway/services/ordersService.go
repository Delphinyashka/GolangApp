package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const ordersServiceURL = "http://localhost:8084" // Orders service URL

// FetchOrders retrieves order data from the orders service based on order IDs
func FetchOrders(orderIDs []string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/orders?ids=%s", ordersServiceURL, strings.Join(orderIDs, ","))
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
