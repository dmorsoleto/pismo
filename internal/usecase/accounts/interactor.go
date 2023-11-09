package accounts

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
)

type AccountsUseCase interface {
	Get(id string) (entity.Account, error)
	Add(account accounts.AddAccount) (string, error)
}
