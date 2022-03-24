package copilot

import (
	"errors"
	"fmt"
)

// below are a list of user events which can be helpful instead of remembering the strings
const (
	EventTypePreexistingSyncStarted         = "preexisting_sync_started"
	EventTypePreexistingSyncCompleted       = "preexisting_sync_completed"
	EventTypePreexistingUserCreated         = "preexisting_user_created"
	EventTypePreexistingThingCreated        = "preexisting_thing_created"
	EventTypePreexistingUserThingAssociated = "preexisting_user_thing_associated"
)

// PreexistingUserEventPayload is a preexisting user
type PreexistingUserEventPayload struct {
	UserID                 *string `json:"user_id,omitempty"`
	FirstName              *string `json:"first_name,omitempty"`
	LastName               *string `json:"last_name,omitempty"`
	Email                  *string `json:"email,omitempty"`
	UTCOffset              *string `json:"utc_offset,omitempty"`
	OriginalCreationDate   *int64  `json:"original_creation_date,omitempty"`
	CopilotAnalysisConsent *bool   `json:"copilot_analysis_consent,omitempty"`
}

// Preexisting is a preexisting thing
type PreexistingThingCreatedPayload struct {
	ThingID              *string `json:"thing_id,omitempty"`
	UserID               *string `json:"user_id,omitempty"`
	FirmwareVersion      *string `json:"firmware_version,omitempty"`
	Model                *string `json:"model,omitempty"`
	OriginalCreationDate *int64  `json:"original_creation_date,omitempty"`
}

// SyncStarted tells Copilot that a sync of preexisting data has started
func SyncStarted(timestamp int64, eventID string) error {
	if eventID == "" {
		eventID = fmt.Sprintf("%s-%d", EventTypePreexistingSyncStarted, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypePreexistingSyncStarted,
		Timestamp: timestamp,
		Payload:   map[string]string{},
	}
	return event.sendEvent()
}

// SyncCompleted tells Copilot that a sync of preexisting data has completed
func SyncCompleted(timestamp int64, eventID string) error {
	if eventID == "" {
		eventID = eventIDHelper(EventTypePreexistingSyncCompleted, "", timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypePreexistingSyncCompleted,
		Timestamp: timestamp,
		Payload:   map[string]string{},
	}
	return event.sendEvent()
}

// PreexistingUserCreated tells Copilot that a user has previously been created. Ideally, the payload.OriginalCreationDate
// field should be set to a Unix timestamp in milliseconds of when the user first was created.
func PreexistingUserCreated(userID string, timestamp int64, eventID string, payload *PreexistingUserEventPayload) error {
	// basic error checking and set some defaults
	if userID == "" {
		return errors.New("userID cannot be blank")
	}

	if payload == nil {
		payload = &PreexistingUserEventPayload{}
	}
	payload.UserID = &userID

	if eventID == "" {
		eventID = eventIDHelper(EventTypePreexistingUserCreated, userID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypePreexistingUserCreated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// PreexistingThingCreated tells Copilot that a thing has previously been created. Ideally, the payload.OriginalCreationDate
// field should be set to a Unix timestamp in milliseconds of when the user first was created.
func PreexistingThingCreated(thingID string, timestamp int64, eventID string, payload *PreexistingThingCreatedPayload) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	if payload == nil {
		payload = &PreexistingThingCreatedPayload{}
	}
	payload.ThingID = &thingID

	if eventID == "" {
		eventID = eventIDHelper(EventTypePreexistingThingCreated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypePreexistingThingCreated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// PreexistingThingUserAssociated tells Copilot about a preexisting thing/user association
func PreexistingThingUserAssociated(thingID string, userID string, timestamp int64, eventID string, originalAssociationDate int64) error {
	// basic error checking and set some defaults
	if thingID == "" || userID == "" {
		return errors.New("thingID and userID cannot be blank")
	}

	payload := map[string]interface{}{
		"user_id":  userID,
		"thing_id": thingID,
	}
	if originalAssociationDate != 0 {
		payload["original_association_date"] = originalAssociationDate
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypePreexistingUserThingAssociated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypePreexistingUserThingAssociated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}
