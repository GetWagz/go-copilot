package copilot

import (
	"errors"
	"fmt"
	"time"
)

// Event is a singular instance of something that you want to collect. The payload will differ depending
// on the event type.
type Event struct {
	Type      string      `json:"type"`
	EventID   string      `json:"event_id"`
	Timestamp int64       `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

func (event *Event) processDefaults() {
	if event.Timestamp == 0 {
		event.Timestamp = time.Now().UnixMilli()
	}
}

// sendEvent takes the event and sends it to copilot, checking for errors; this consolidates
// the general collect event call checks
func (event *Event) sendEvent() error {
	event.processDefaults()

	eventRequest := eventRequest{
		Events: []Event{
			*event,
		},
	}
	response, eventError, err := makeCollectAPICall(eventRequest)
	if err != nil {
		return err
	}
	if eventError != nil {
		return eventError
	}
	if response == nil {
		return errors.New("invalid client request")
	}
	if len(response.InvalidEvents) > 0 {
		// find the corresponding event error and return it
		// it should usually only be one, but in case they somehow got batched, it's
		// best to verify
		for _, ie := range response.InvalidEvents {
			if ie.EventID == event.EventID {
				return &ie
			}
		}
	}
	return nil
}

// eventRequest is the request that is sent to the collect API
type eventRequest struct {
	Events []Event `json:"events"`
}

// EventResponse is the response to a call. Since Copilot sends back a 200 on events,
// the best way to tell if there was an error is to verify the length of this
// InvalidEvents slice is 0
type EventResponse struct {
	InvalidEvents []InvalidEventError `json:"invalid_events"`
}

// InvalidEventError represents an invalid event error returned from Copilot's collect API
type InvalidEventError struct {
	EventID    string `json:"event_id"`
	Index      int    `json:"index"`
	EventError string `json:"error"`
}

func (err *InvalidEventError) Error() string {
	return fmt.Sprintf("%s-%s", err.EventID, err.EventError)
}

// EventResponseError represents an error from the Copilot Collect API that does not directly
// related to a single event. Examples include missing authentication or missing fields.
type EventResponseError struct {
	ErrorMessage string `json:"error_message"`
	Reason       string `json:"reason"`
}

func (err *EventResponseError) Error() string {
	return fmt.Sprintf("%s-%s", err.Reason, err.ErrorMessage)
}

func eventIDHelper(eventType, key string, timestamp int64) string {
	if timestamp == 0 {
		timestamp = time.Now().UnixMilli()
	}
	id := fmt.Sprintf("%s-%s-%d", eventType, key, timestamp)

	// check if we need to trim it
	if len(id) > 50 {
		id = id[0:50]
	}
	return id
}
