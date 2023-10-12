package schema

type ResponseSpendingType struct {
	ID                 string `json:"id"`
	ProfileID          string `json:"profile_id"`
	Title              string `json:"title"`
	MaximumLimit       int    `json:"maximum_limit"`
	FormatMaximumLimit string `json:"format_maximum_limit"`
	Icon               string `json:"icon"`
}

type ResponseSpendingTypeJoinTable struct {
	ID                 string `json:"id"`
	ProfileID          string `json:"profile_id"`
	Title              string `json:"title"`
	MaximumLimit       int    `json:"maximum_limit"`
	FormatMaximumLimit string `json:"format_maximum_limit"`
	Icon               string `json:"icon"`
	Used               int    `json:"used"`
	FormatUsed         string `json:"format_used"`
	PersentaseUsed     string `json:"persentase_used"`
}
