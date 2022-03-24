package copilot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Printf("\n-----------------------\n%v\n%v\n", response.StatusCode, err)

	if response.StatusCode != http.StatusOK {
		// parse the error message and return
		errorResponse := &EventResponseError{}
		err = json.NewDecoder(response.Body).Decode(errorResponse)
		fmt.Printf("\nError response\n%+v\n", errorResponse)
		return nil, errorResponse, err
	}
	// Copilot returns a 200 even if there are invalid events, so we need
	// to determine if there are any invalid events and then return them

	eventResponse := &EventResponse{}
	err = json.NewDecoder(response.Body).Decode(eventResponse)
	fmt.Printf("\nEvent response\n%+v\n", eventResponse)
	return eventResponse, nil, err
}
