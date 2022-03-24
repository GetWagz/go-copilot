package copilot

import (
	"errors"
	"fmt"
)

// below are a list of user events which can be helpful instead of remembering the strings
const (
	EventTypeUserCreated = "user_created"
	EventTypeUserUpdated = "user_updated"
)

// UserEventPayload is used to set data for both the UserCreated and UserUpdated calls
type UserEventPayload struct {
	UserID    *string `json:"user_id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	UTCOffset *string `json:"utc_offset"`
}

// UserCreated tells copilot a new user was created. The userID is required.
// All other fields can be blank and a sane default will be used.
func UserCreated(userID string, timestamp int64, eventID string, payload *UserEventPayload) error {
	// basic error checking and set some defaults
	if userID == "" {
		return errors.New("userID cannot be blank")
	}

	if payload == nil {
		payload = &UserEventPayload{}
	}
	payload.UserID = &userID

	if eventID == "" {
		eventID = fmt.Sprintf("%s-%s-%d", EventTypeUserCreated, userID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeUserCreated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// UserUpdated updates the user in copilot's system. The userID is required.
// All other fields can be blank and a sane default will be used.
func UserUpdated(userID string, timestamp int64, eventID string, payload *UserEventPayload) error {
	// basic error checking and set some defaults
	if userID == "" {
		return errors.New("userID cannot be blank")
	}

	if payload == nil {
		payload = &UserEventPayload{}
	}
	payload.UserID = &userID

	if eventID == "" {
		eventID = fmt.Sprintf("%s-%s-%d", EventTypeUserUpdated, userID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeUserUpdated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}
