package copilot

import (
	"errors"
	"log"
	"os"
)

// hold the configuration
type configStruct struct {
	ClientID        string
	ClientSecret    string
	CollectEndpoint string
	ConsentEndpoint string
}

var config *configStruct = nil

func init() {
	// we call directly into setup; we do it this way so the user
	// can either have the client configured by the environment or they
	// can pass in the values directly
	clientID := osHelper("COPILOT_CLIENT_ID", "")
	clientSecret := osHelper("COPILOT_CLIENT_SECRET", "")
	collectEndpoint := osHelper("COPILOT_CLIENT_COLLECT_ENDPOINT", "")
	consentEndpoint := osHelper("COPILOT_CLIENT_CONSENT_ENDPOINT", "")
	Setup(clientID, clientSecret, collectEndpoint, consentEndpoint)
}

// Setup is called on init from the environment but can also be called explicitly
// by tests or the client
func Setup(clientID string, clientSecret string, collectEndpoint, consentEndpoint string) error {
	// if they are missing, we want to log an error but we shouldn't
	// nuke the caller through a panic
	if clientID == "" || clientSecret == "" || collectEndpoint == "" {
		message := "copilot requires the client credentials and endpoint to be configured; no calls will be processed"
		log.Print(message)
		return errors.New(message)
	}

	config = &configStruct{
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		CollectEndpoint: collectEndpoint,
		ConsentEndpoint: consentEndpoint,
	}
	return nil
}

// IsSetUp is a helper to determine if the copilot client is configured. Note that this
// does not determin if it is set up correctly or that credentials are valid!
func IsSetUp() bool {
	return config != nil
}

// osHelper provides a quick and easy way to get defaults from the environment
func osHelper(key, defaultValue string) string {
	found := os.Getenv(key)
	if found == "" {
		found = defaultValue
	}
	return found
}

// String takes a string and converts it to a pointer to be used
// in the parameter passing
func String(str string) *string {
	return &str
}

// Int64 takes an int64 and converts it to a pointer to be used
// in the parameter passing, most often for timestamps
func Int64(i int64) *int64 {
	return &i
}

// Bool takes a bool and converts it to a pointer to be used
// in the parameter passing
func Bool(b bool) *bool {
	return &b
}
