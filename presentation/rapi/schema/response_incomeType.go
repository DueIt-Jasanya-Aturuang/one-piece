package schema

type ResponseIncomeType struct {
	ID          string  `json:"id"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
}
