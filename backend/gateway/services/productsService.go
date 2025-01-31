package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const productsServiceURL = "http://localhost:8085"

func FetchProductsBatch(productIDs map[string]bool) ([]map[string]interface{}, error) {
	var ids []string
	for id := range productIDs {
		ids = append(ids, id)
	}

	url := fmt.Sprintf("%s/products?ids=%s", productsServiceURL, strings.Join(ids, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from products service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch products: status code %d", resp.StatusCode)
	}

	var products []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return products, nil
}
