package copilot

import (
	"errors"
)

// below is custom event constant
const (
	EventTypeCustomEvent = "custom_event"
)

type CustomEventPayload map[string]interface{}

// CustomEvent represents a single custom event sent to copilot. It must have a set subtype. In the payload
// either a user_id or thing_id string must be provided. Both can be provided. All other keys on the
// payload will be sent as is.
func CustomEvent(eventSubtype string, timestamp int64, eventID string, payload CustomEventPayload) error {
	if eventSubtype == "" {
		return errors.New("you must provide a subtype")
	}
	if payload == nil {
		payload = CustomEventPayload{}
	}
	payload["subtype"] = eventSubtype
	// the payload must contain either a user_id or a thing_id
	_, foundUser := payload["user_id"]
	_, foundThing := payload["thing_id"]
	if !foundUser && !foundThing {
		return errors.New("either a user_id or a thing_id must be included in the payload")
	}
	if eventID == "" {
		eventID = eventIDHelper(EventTypeCustomEvent, eventSubtype, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeCustomEvent,
		Timestamp: timestamp,
		Payload:   payload,
	}

	return event.sendEvent()
}
