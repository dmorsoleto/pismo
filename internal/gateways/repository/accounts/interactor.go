package accounts

import "dmorsoleto/internal/entity"

type AccountsRepository interface {
	Get(id string) (entity.Account, error)
	Add(account AddAccount) (string, error)
}
