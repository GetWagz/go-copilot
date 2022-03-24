package copilot

type ConsentRequest struct {
	UserID       string `json:"user_id"`
	ConsentValue bool   `json:"consent_value"`
}
