package copilot

import (
	"errors"
	"fmt"
)

const (
	EventTypeUnsubscribe = "unsubscribe"
)

// UnsubscribeUserEmail tells Copilot that a user has unsubscribed from email notifications. This should be sent
// when the user unsubscribes from emails on your system.
func UnsubscribeUserEmail(email string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if email == "" {
		return errors.New("email cannot be blank")
	}

	payload := map[string]string{
		"email": email,
	}

	if eventID == "" {
		eventID = fmt.Sprintf("%s-%s-%d", EventTypeUnsubscribe, email, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeUnsubscribe,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}
