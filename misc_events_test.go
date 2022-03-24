package copilot_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/GetWagz/go-copilot"
	"github.com/stretchr/testify/assert"
)

func TestUnsubscribeEvent(t *testing.T) {
	email := fmt.Sprintf("%d@wagz.com", rand.Int63n(9999999))
	timestamp := time.Now().UnixMilli()

	// make sure the userID is set
	err := copilot.UnsubscribeUserEmail("", timestamp, "")
	assert.NotNil(t, err)

	if !copilot.IsSetUp() {
		t.SkipNow()
	}
	err = copilot.UnsubscribeUserEmail(email, timestamp, "")
	assert.Nil(t, err)

}
