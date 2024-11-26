package repository

import (
	"errors"
	"example/entity"
	"example/internal/common/helper/sqlormhelper"
	"example/internal/dto"
	"gorm.io/gorm"
)

type (
	AccountRepository interface {
		FindOneAccount(input *dto.FindOneAccountInput) (*entity.Account, error)
		FindOneAccountTx(tx *gorm.DB, input *dto.FindOneAccountInput) (*entity.Account, error)
		TransferTx(fromAcct string, toAcct string, amount float64, fn func(transferTxFuncInput *dto.TransferTxFuncInput) (bool, error)) error
	}

	accountRepository struct {
		postgresOrmDb sqlormhelper.SqlGormDatabase
	}
)

func NewAccountRepository(postgresOrmDb sqlormhelper.SqlGormDatabase) AccountRepository {
	return &accountRepository{postgresOrmDb: postgresOrmDb}
}

func (a *accountRepository) FindOneAccount(input *dto.FindOneAccountInput) (*entity.Account, error) {
	db := a.postgresOrmDb.Open()

	account := &entity.Account{}
	accountCond := &entity.Account{
		AccountNo: input.AccountNo,
	}

	result := db.Where(accountCond).First(account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func (a *accountRepository) FindOneAccountTx(tx *gorm.DB, input *dto.FindOneAccountInput) (*entity.Account, error) {
	account := &entity.Account{}
	accountCond := &entity.Account{
		AccountNo: input.AccountNo,
	}

	result := tx.Where(accountCond).First(account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func (a *accountRepository) TransferTx(fromAcctNo string, toAcctNo string, amount float64,
	fn func(transferTxFuncInput *dto.TransferTxFuncInput) (bool, error)) error {
	return runInTx(a.postgresOrmDb.Open(), func(db *gorm.DB) error {
		fromAcct, err := a.FindOneAccountTx(db, &dto.FindOneAccountInput{AccountNo: fromAcctNo})
		if err != nil {
			return err
		}

		toAcct, err := a.FindOneAccountTx(db, &dto.FindOneAccountInput{AccountNo: toAcctNo})
		if err != nil {
			return err
		}

		transferTxFuncInput := &dto.TransferTxFuncInput{
			FromAcct: fromAcct,
			ToAcct:   toAcct,
			Amount:   amount,
		}
		result, err := fn(transferTxFuncInput)
		if err != nil {
			return err
		}
		if !result {
			return errors.New("transfer failed")
		}

		rs := db.Model(&entity.Account{}).Where(&entity.Account{AccountNo: fromAcctNo}).Update(
			"AccountBalance", transferTxFuncInput.FromAcct.AccountBalance,
		)
		if rs.Error != nil {
			return rs.Error
		}

		rs = db.Model(&entity.Account{}).Where(&entity.Account{AccountNo: toAcctNo}).Update(
			"AccountBalance", transferTxFuncInput.ToAcct.AccountBalance,
		)
		if rs.Error != nil {
			return rs.Error
		}

		return nil
	})
}
