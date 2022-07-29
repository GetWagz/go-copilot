package copilot

import (
	"errors"
)

// below are a list of user events which can be helpful instead of remembering the strings
const (
	EventTypeUserCreated = "user_created"
	EventTypeUserUpdated = "user_updated"
	EventTypeUserDeleted = "user_deleted"
)

// UserEventPayload is used to set data for both the UserCreated and UserUpdated calls
type UserEventPayload struct {
	UserID    *string `json:"user_id,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	UTCOffset *string `json:"utc_offset,omitempty"`
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
		eventID = eventIDHelper(EventTypeUserCreated, userID, timestamp)
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
		eventID = eventIDHelper(EventTypeUserUpdated, userID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeUserUpdated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// UserDeleted tells copilot that a user has been deleted
func UserDeleted(userID string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if userID == "" {
		return errors.New("userID cannot be blank")
	}

	payload := &UserEventPayload{
		UserID: &userID,
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeUserDeleted, userID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeUserDeleted,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}
