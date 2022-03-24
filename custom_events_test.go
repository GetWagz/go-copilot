package copilot_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/GetWagz/go-copilot"
	"github.com/stretchr/testify/assert"
)

func TestCustomEvents(t *testing.T) {
	userID := fmt.Sprintf("test---%d", rand.Int63n(9999999))
	thingID := fmt.Sprintf("test---%d", rand.Int63n(9999999))
	email := fmt.Sprintf("%s@wagz.com", userID)
	timestamp := time.Now().UnixMilli()

	// make sure the required fields are checked
	err := copilot.CustomEvent("", timestamp, "", nil)
	assert.NotNil(t, err)
	err = copilot.CustomEvent("subtype", timestamp, "", nil)
	assert.NotNil(t, err)
	additionalParams := copilot.CustomEventPayload{}
	err = copilot.CustomEvent("subtype", timestamp, "", additionalParams)
	assert.NotNil(t, err)
	additionalParams["sample"] = 3
	err = copilot.CustomEvent("subtype", timestamp, "", additionalParams)
	assert.NotNil(t, err)

	if !copilot.IsSetUp() {
		t.SkipNow()
	}

	additionalParams["user_id"] = userID
	additionalParams["thing_id"] = thingID
	additionalParams["email"] = email
	err = copilot.CustomEvent("subtype", timestamp, "", additionalParams)
	assert.Nil(t, err)

}
