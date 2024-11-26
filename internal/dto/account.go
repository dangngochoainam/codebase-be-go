package dto

import "example/entity"

type TransferMoneyRequestDTO struct {
	FromAcctNo string  `json:"fromAcctNo"`
	ToAcctNo   string  `json:"toAcctNo"`
	Amount     float64 `json:"amount"`
}

// Input - Output Repository
type FindOneAccountInput struct {
	AccountNo string
}

type TransferTxFuncInput struct {
	FromAcct *entity.Account
	ToAcct   *entity.Account
	Amount   float64
}
