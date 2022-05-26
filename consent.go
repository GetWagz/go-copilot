package copilot

// UpdateUserConsent updates the user's consent using the consent endpoint
// https://docs.copilot.cx/docs/server-api-your-own/reference/consent-api-reference
func UpdateUserConsent(userID string, consentValue bool) error {
	return makeConsentCall(userID, consentValue)
}
