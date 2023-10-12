package schema

type ResponseBalance struct {
	ID                        string `json:"id"`
	ProfileID                 string `json:"profile_id"`
	TotalIncomeAmount         int    `json:"total_income_amount"`
	TotalIncomeAmountFormat   string `json:"total_income_amount_format"`
	TotalSpendingAmount       int    `json:"total_spending_amount"`
	TotalSpendingAmountFormat string `json:"total_spending_amount_format"`
	Balance                   int    `json:"balance"`
	BalanceFormat             string `json:"balance_format"`
}
