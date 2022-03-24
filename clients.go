package copilot

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var httpClient = http.Client{Timeout: 5 * time.Second}

func makeCollectAPICall(data eventRequest) (*EventResponse, *EventResponseError, error) {
	if config == nil {
		return nil, nil, errors.New("copilot client not configured")
	}
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	httpBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, config.CollectEndpoint, httpBody)
	if err != nil {
		return nil, nil, err
	}
	req.SetBasicAuth(config.ClientID, config.ClientSecret)
	req.Header.Add("content-type", "application/json")

	// now make the call
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// parse the error message and return
		errorResponse := &EventResponseError{}
		err = json.NewDecoder(response.Body).Decode(errorResponse)
		return nil, errorResponse, err
	}
	// Copilot returns a 200 even if there are invalid events, so we need
	// to determine if there are any invalid events and then return them

	eventResponse := &EventResponse{}
	err = json.NewDecoder(response.Body).Decode(eventResponse)
	return eventResponse, nil, err
}
