package service

import "natthan.com/go-play/repository"

func convertAccountResponse(account *repository.Account) AccountResponse {
	return AccountResponse{
		AccountID:   account.AccountID,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	}
}
