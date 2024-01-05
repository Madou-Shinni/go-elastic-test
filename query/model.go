package query

type User struct {
	AccountNumber int    `json:"account_number,omitempty"`
	Balance       int    `json:"balance,omitempty"`
	Address       string `json:"address,omitempty"`
}
