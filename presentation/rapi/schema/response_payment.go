package schema

type ResponsePayment struct {
	ID          string  `json:"id"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}
