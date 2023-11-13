package availablecreditlimit

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
	availablecreditlimit "dmorsoleto/internal/gateways/repository/availableCreditLimit"
	"dmorsoleto/internal/helpers/uuid"
	"errors"

	"github.com/sirupsen/logrus"
)

type availableCreditLimitUseCase struct {
	availableCreditLimitRepo availablecreditlimit.AvailableCreditLimit
	accountsRepo             accounts.AccountsRepository
	uuidHelper               uuid.UuidHelper
}

func NewAvailableCreditLimitUseCase(availableCreditLimitRepo availablecreditlimit.AvailableCreditLimit, accountsRepo accounts.AccountsRepository, uuidHelper uuid.UuidHelper) AvailableCreditLimit {
	return &availableCreditLimitUseCase{
		availableCreditLimitRepo: availableCreditLimitRepo,
		uuidHelper:               uuidHelper,
		accountsRepo:             accountsRepo,
	}
}

func (ref *availableCreditLimitUseCase) Get(accountId string) (entity.AvailableCreditLimit, error) {
	if ref.uuidHelper.IsValidUUID(accountId) {
		return entity.AvailableCreditLimit{}, errors.New("invalid account id")
	}

	hasCredit, err := ref.availableCreditLimitRepo.Get(accountId)
	if err != nil {
		return entity.AvailableCreditLimit{}, err
	}

	return hasCredit, nil
}

func (ref *availableCreditLimitUseCase) Add(addCredit entity.AddAvailableCreditLimit) (string, error) {
	if !ref.uuidHelper.IsValidUUID(addCredit.AccountId) {
		return "", errors.New("invalid account id")
	}

	hasAccount, err := ref.accountsRepo.Get(addCredit.AccountId)
	if err != nil || hasAccount == (entity.Account{}) {
		logrus.Error(err)
		return "", errors.New("account not found")
	}

	if addCredit.AvailableCreditLimit <= 0 {
		return "", errors.New("invalid available credit limit")
	}

	lastID, err := ref.availableCreditLimitRepo.Add(addCredit)
	if err != nil {
		return "", err
	}

	return lastID, nil
}
