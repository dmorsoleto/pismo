package availablecreditlimit

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/helpers/database"
	"dmorsoleto/internal/helpers/uuid"

	"github.com/jmoiron/sqlx"
)

const (
	selectCreditAvailableByAccountID = "SELECT * FROM pismo.available_credit_limit WHERE account_id = $1"
	addCreditAvailable               = "INSERT INTO pismo.availablecreditlimit (available_credit_limit_id, account_id, available_credit_limit) VALUES($1, $2, $3)"
)

type availableCreditLimit struct {
	db             *sqlx.DB
	databaseHelper database.DatabaseHelper
	uuid           uuid.UuidHelper
}

func NewAvailableCreditLimit(databaseHelper database.DatabaseHelper, uuid uuid.UuidHelper) AvailableCreditLimit {
	db := databaseHelper.GetConnection()

	return &availableCreditLimit{
		db:   db,
		uuid: uuid,
	}
}

func (ref *availableCreditLimit) Get(accountId string) (entity.AvailableCreditLimit, error) {
	credit := entity.AvailableCreditLimit{}

	err := ref.db.Get(&credit, selectCreditAvailableByAccountID, accountId)
	if err != nil {
		return entity.AvailableCreditLimit{}, err
	}

	return credit, nil
}

func (ref *availableCreditLimit) Add(addCredit entity.AddAvailableCreditLimit) (string, error) {

	id := ref.uuid.Generate()

	_, err := ref.db.Exec(addCreditAvailable, id, addCredit.AccountId, addCredit.AvailableCreditLimit)
	if err != nil {
		return "", err
	}

	return id, nil
}
