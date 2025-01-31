package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const clientsServiceURL = "http://localhost:8083"

func FetchClientsBatch(clientIDs map[string]bool) ([]map[string]interface{}, error) {
	var ids []string
	for id := range clientIDs {
		ids = append(ids, id)
	}

	url := fmt.Sprintf("%s/clients?ids=%s", clientsServiceURL, strings.Join(ids, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients from clients service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch clients: status code %d", resp.StatusCode)
	}

	var clients []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	err = json.Unmarshal(body, &clients)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return clients, nil
}
