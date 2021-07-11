package service

import (
	"strings"
	"time"

	"github.com/Mersock/golang-hexagonal-architecture/errs"
	"github.com/Mersock/golang-hexagonal-architecture/logs"
	"github.com/Mersock/golang-hexagonal-architecture/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(custId int, req NewAccountReq) (*AccountRes, error) {
	//validate
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amout at least 5000")
	}

	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  custId,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}
	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	res := AccountRes{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &res, nil
}

func (s accountService) GetAccounts(custId int) ([]AccountRes, error) {
	accounts, err := s.accRepo.GetAll(custId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	res := []AccountRes{}

	for _, account := range accounts {
		res = append(res, AccountRes{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return res, nil
}
