package usecase

import (
	"example/internal/common/helper/copyhepler"
	"example/internal/dto"
	"example/internal/repository"
)

type (
	AccountUseCase interface {
		TransferMoney(input *dto.TransferMoneyRequestDTO) error
	}

	accountUseCase struct {
		accountRepository repository.AccountRepository
		modelConverter    copyhepler.ModelConverter
	}
)

func NewAccountUseCase(accountRepository repository.AccountRepository,
	modelConverter copyhepler.ModelConverter) AccountUseCase {
	return &accountUseCase{
		accountRepository: accountRepository,
		modelConverter:    modelConverter,
	}
}

func (u *accountUseCase) TransferMoney(input *dto.TransferMoneyRequestDTO) error {
	return u.accountRepository.TransferTx(input.FromAcctNo, input.ToAcctNo, input.Amount, func(input *dto.TransferTxFuncInput) (bool, error) {
		if input.FromAcct.AccountBalance > input.Amount {
			input.FromAcct.AccountBalance -= input.Amount
			input.ToAcct.AccountBalance += input.Amount
			return true, nil
		}
		return false, nil
	})
}
