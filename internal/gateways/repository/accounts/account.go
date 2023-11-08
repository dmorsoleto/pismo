package accounts

import (
	"dmorsoleto/internal/entity"

	"github.com/jmoiron/sqlx"
)

const (
	selectAccountByID = "SELECT * FROM accounts WHERE account_ID = $1"
	insertAccount     = "INSERT INTO accounts (document_number) VALUES (:document_number)"
)

type accountsRepository struct {
	db *sqlx.DB
}

func NewAccountsRepository(db *sqlx.DB) AccountsRepository {
	return &accountsRepository{db: db}
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

	_, err := ref.db.NamedExec(insertAccount, account)

	if err != nil {
		return err
	}

	return nil
}
