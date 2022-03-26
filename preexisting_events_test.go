package copilot_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/GetWagz/go-copilot"
	"github.com/stretchr/testify/assert"
)

func TestPreExistingSyncCalls(t *testing.T) {
	userID := fmt.Sprintf("test---%d", rand.Int63n(9999999))
	email := fmt.Sprintf("%s@wagz.com", userID)
	thingID := fmt.Sprintf("%d", rand.Int63n(99999999))
	timestamp := time.Now().UnixMilli()
	originalCreation := time.Now().AddDate(0, 0, -50).UnixMilli()

	// do some very basic error checking
	// since the sync start and complete calls don't have any
	// required fields, we can't verify the calls locally

	err := copilot.PreexistingThingCreated("", timestamp, "", nil)
	assert.NotNil(t, err)
	err = copilot.PreexistingUserCreated("", timestamp, "", nil)
	assert.NotNil(t, err)
	err = copilot.PreexistingThingUserAssociated("", "", timestamp, "", 0)
	assert.NotNil(t, err)

	if !copilot.IsSetUp() {
		t.SkipNow()
	}

	err = copilot.SyncStarted(0, "")
	assert.Nil(t, err)

	userParams := &copilot.PreexistingUserEventPayload{
		Email:                copilot.String(email),
		OriginalCreationDate: copilot.Int64(originalCreation),
	}
	err = copilot.PreexistingUserCreated(userID, timestamp, "", userParams)
	assert.Nil(t, err)

	thingParams := &copilot.PreexistingThingCreatedPayload{
		Model:                copilot.String("model"),
		OriginalCreationDate: copilot.Int64(originalCreation),
	}
	err = copilot.PreexistingThingCreated(thingID, timestamp, "", thingParams)
	assert.Nil(t, err)

	err = copilot.PreexistingThingUserAssociated(thingID, userID, timestamp, "", originalCreation)
	assert.Nil(t, err)

	err = copilot.SyncCompleted(0, "")
	assert.Nil(t, err)
}
