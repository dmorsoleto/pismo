package accounts

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/helpers/database"
	"dmorsoleto/internal/helpers/uuid"

	"github.com/jmoiron/sqlx"
)

const (
	selectAccountByID = "SELECT * FROM pismo.accounts WHERE account_ID = $1"
	insertAccount     = "INSERT INTO pismo.accounts (account_ID, document_number) VALUES ($1, $2)"
)

type accountsRepository struct {
	db *sqlx.DB

	uuidHelper uuid.UuidHelper
}

func NewAccountsRepository(databaseHelper database.DatabaseHelper, uuidHelper uuid.UuidHelper) AccountsRepository {
	db := databaseHelper.GetConnection()

	return &accountsRepository{db: db, uuidHelper: uuidHelper}
}

func (ref *accountsRepository) Get(id string) (entity.Account, error) {
	account := entity.Account{}
	err := ref.db.Get(&account, selectAccountByID, id)

	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (ref *accountsRepository) Add(account AddAccount) error {

	id := ref.uuidHelper.Generate()
	_, err := ref.db.Exec(insertAccount, id, account.DocumentNumber)

	if err != nil {
		return err
	}

	return nil
}
