package copilot

import (
	"errors"
)

// below are a list of thing events which can be helpful instead of remembering the strings
const (
	EventTypeThingCreated                  = "thing_created"
	EventTypeThingUpdated                  = "thing_updated"
	EventTypeThingAssociated               = "user_thing_associated"
	EventTypeThingDisassociated            = "user_thing_disassociated"
	EventTypeThingStatusChanged            = "thing_status_changed"
	EventTypeThingInteraction              = "thing_interaction"
	EventTypeThingConnected                = "thing_connected"
	EventTypeThingConsumableUsage          = "consumable_usage"
	EventTypeThingFirmwareUpgradeStarted   = "firmware_upgrade_started"
	EventTypeThingFirmwareUpgradeCompleted = "firmware_upgrade_completed"
)

// ThingCreatedUpdatedPayload is used to set data for both the ThingCreated and ThingUpdated calls
type ThingCreatedUpdatedPayload struct {
	ThingID         *string `json:"thing_id,omitempty"`
	UserID          *string `json:"user_id,omitempty"`
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	Model           *string `json:"model,omitempty"`
}

// ThingStatusChangedPayload tells Copilot that the status of the thing has changed. This can include
// things like activation status, battery level, connectivity level, or other changes to the thing's state.
type ThingStatusChangedPayload struct {
	ThingID     *string `json:"thing_id,omitempty"`
	StatusKey   *string `json:"status_key,omitempty"`
	StatusValue *string `json:"status_value,omitempty"`
	StatusDate  *int64  `json:"status_date,omitempty"`
	UserID      *string `json:"user_id,omitempty"`
}

// ThingInteractionEventPayload holds arbitrary key/values sent as part of the Thing Interaction endpoint call.
type ThingInteractionEventPayload map[string]interface{}

// ThingCreated tells Copilot a thing has been created. The thingID is required.
// All other fields can be blank and a sane default will be used.
func ThingCreated(thingID string, timestamp int64, eventID string, payload *ThingCreatedUpdatedPayload) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	if payload == nil {
		payload = &ThingCreatedUpdatedPayload{}
	}
	payload.ThingID = &thingID

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingCreated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingCreated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingUpdated tells Copilot a thing has been updated. The thingID is required.
// All other fields can be blank and a sane default will be used.
func ThingUpdated(thingID string, timestamp int64, eventID string, payload *ThingCreatedUpdatedPayload) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	if payload == nil {
		payload = &ThingCreatedUpdatedPayload{}
	}
	payload.ThingID = &thingID

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingUpdated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingUpdated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingAssociated tells Copilot that a thing has been associated to a user
func ThingAssociated(thingID string, userID string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" || userID == "" {
		return errors.New("thingID and userID cannot be blank")
	}

	payload := map[string]string{
		"user_id":  userID,
		"thing_id": thingID,
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingAssociated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingAssociated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingDisassociated tells Copilot that a thing has been disassociated from a user
func ThingDisassociated(thingID string, userID string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" || userID == "" {
		return errors.New("thingID and userID cannot be blank")
	}

	payload := map[string]string{
		"user_id":  userID,
		"thing_id": thingID,
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingDisassociated, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingDisassociated,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingStatusChanged tells Copilot that the status of the thing has changed (see the comments on the ThingStatusChangedPayload).
// Note that if the StatusDate field is nil or 0, we will set it to the timestamp's value. The only payload field that is not
// required is the userID.
func ThingStatusChanged(thingID string, timestamp int64, eventID string, payload *ThingStatusChangedPayload) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	if payload == nil {
		return errors.New("payload is required")
	}

	if payload.StatusKey == nil || payload.StatusValue == nil {
		return errors.New("StatusKey and StatusValue are required and cannot be blank")
	}
	if *payload.StatusKey == "" || *payload.StatusValue == "" {
		return errors.New("StatusKey and StatusValue are required and cannot be blank")
	}

	if payload.StatusDate == nil || *payload.StatusDate == 0 {
		payload.StatusDate = &timestamp
	}

	payload.ThingID = &thingID

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingStatusChanged, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingStatusChanged,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingIneraction tells Copilot about an arbitrary interaction. You can set any fields you want in the payload and they
// will be passed straight through.
func ThingIneraction(thingID string, timestamp int64, eventID string, payload ThingInteractionEventPayload) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	if payload == nil {
		payload = ThingInteractionEventPayload{}
	}
	payload["thing_id"] = thingID

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingInteraction, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingInteraction,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingConnected tells Copilot that a thing has been connected
func ThingConnected(thingID string, userID string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	payload := map[string]string{
		"thing_id": thingID,
	}
	if userID != "" {
		payload["user_id"] = userID
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingConnected, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingConnected,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingConsumableUsage tells Copilot that the thing consumed something. For example, if the thing is a
// printer and prints a sheet of paper, this function could be called to tell Copilot that the thing
// consumed paper.
func ThingConsumableUsage(thingID string, userID string, consumableType string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	payload := map[string]string{
		"thing_id": thingID,
	}
	if userID != "" {
		payload["user_id"] = userID
	}
	if consumableType != "" {
		payload["consumable_type"] = consumableType
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingConsumableUsage, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingConsumableUsage,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingFirmwareUpgradeStarted tells Copilot that a firmware upgrade has begin on the thing.
func ThingFirmwareUpgradeStarted(thingID string, userID string, firmwareVersion string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	payload := map[string]string{
		"thing_id": thingID,
	}
	if userID != "" {
		payload["user_id"] = userID
	}
	if firmwareVersion != "" {
		payload["firmware_version"] = firmwareVersion
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingFirmwareUpgradeStarted, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingFirmwareUpgradeStarted,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}

// ThingFirmwareUpgradeCompleted tells Copilot that a firmware upgrade has completed on the thing.
func ThingFirmwareUpgradeCompleted(thingID string, userID string, firmwareVersion string, timestamp int64, eventID string) error {
	// basic error checking and set some defaults
	if thingID == "" {
		return errors.New("thingID cannot be blank")
	}

	payload := map[string]string{
		"thing_id": thingID,
	}
	if userID != "" {
		payload["user_id"] = userID
	}
	if firmwareVersion != "" {
		payload["firmware_version"] = firmwareVersion
	}

	if eventID == "" {
		eventID = eventIDHelper(EventTypeThingFirmwareUpgradeCompleted, thingID, timestamp)
	}

	event := Event{
		EventID:   eventID,
		Type:      EventTypeThingFirmwareUpgradeCompleted,
		Timestamp: timestamp,
		Payload:   payload,
	}
	return event.sendEvent()
}
