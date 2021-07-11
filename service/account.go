package service

type NewAccountReq struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountRes struct {
	AccountID   int     `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountReq) (*AccountRes, error)
	GetAccounts(int) ([]AccountRes, error)
}
