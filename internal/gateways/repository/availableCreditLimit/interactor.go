package availablecreditlimit

import "dmorsoleto/internal/entity"

type AvailableCreditLimit interface {
	Get(accountId string) (entity.AvailableCreditLimit, error)
	Add(addCredit entity.AddAvailableCreditLimit) (string, error)
}
