package copilot_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/GetWagz/go-copilot"
	"github.com/stretchr/testify/assert"
)

func TestUserEvents(t *testing.T) {
	userID := fmt.Sprintf("test---%d", rand.Int63n(9999999))
	email := fmt.Sprintf("%s@wagz.com", userID)
	timestamp := time.Now().UnixMilli()

	// make sure the userID is set
	err := copilot.UserCreated("", timestamp, "", nil)
	assert.NotNil(t, err)
	err = copilot.UserUpdated("", timestamp, "", nil)
	assert.NotNil(t, err)

	if !copilot.IsSetUp() {
		t.SkipNow()
	}
	additionalParams := copilot.UserEventPayload{
		Email:     copilot.String(email),
		FirstName: copilot.String("Test"),
		LastName:  copilot.String("Test"),
		UTCOffset: copilot.String("-0500"),
	}
	err = copilot.UserCreated(userID, timestamp, "", &additionalParams)
	assert.Nil(t, err)
	timestamp += 100
	err = copilot.UserUpdated(userID, timestamp, "", &additionalParams)
	assert.Nil(t, err)

}
