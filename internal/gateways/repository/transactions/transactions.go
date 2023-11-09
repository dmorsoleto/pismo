package transactions

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/helpers/database"
	"dmorsoleto/internal/helpers/uuid"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	insertTransaction = "INSERT INTO pismo.transactions (transaction_id, account_id, operations_type_id, amount, event_date) VALUES ($1, $2)"
)

type transactionsRepository struct {
	db         *sqlx.DB
	uuidHelper uuid.UuidHelper
}

func NewTransactionsRepository(databaseHelper database.DatabaseHelper, uuidHelper uuid.UuidHelper) TransactionsRepository {
	db := databaseHelper.GetConnection()

	return &transactionsRepository{
		db:         db,
		uuidHelper: uuidHelper,
	}
}

func NewTransactionsRepositoryFromDB(db *sqlx.DB, uuidHelper uuid.UuidHelper) TransactionsRepository {
	return &transactionsRepository{db: db, uuidHelper: uuidHelper}
}

func (ref *transactionsRepository) Add(transaction entity.Transactions) (string, error) {

	id := ref.uuidHelper.Generate()

	created_time := time.Now()

	_, err := ref.db.Exec(insertTransaction, id, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, created_time)

	if err != nil {
		logrus.Error("Error on insert a transaction", err)
		return "", err
	}

	return id, nil
}
