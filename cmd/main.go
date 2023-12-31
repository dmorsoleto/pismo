package main

import (
	http_controller "dmorsoleto/internal/controller"
	handlers_accounts "dmorsoleto/internal/controller/handlers/accounts"
	handlers_credit "dmorsoleto/internal/controller/handlers/availableCreditLimit"
	handlers_transactions "dmorsoleto/internal/controller/handlers/transactions"
	"dmorsoleto/internal/entity"
	accounts_repository "dmorsoleto/internal/gateways/repository/accounts"
	availablecreditlimit_repo "dmorsoleto/internal/gateways/repository/availableCreditLimit"
	operationstype "dmorsoleto/internal/gateways/repository/operationsType"
	transactions_repository "dmorsoleto/internal/gateways/repository/transactions"
	"dmorsoleto/internal/helpers/database"
	"dmorsoleto/internal/helpers/uuid"
	"dmorsoleto/internal/usecase/accounts"
	availablecreditlimit "dmorsoleto/internal/usecase/availableCreditLimit"
	"dmorsoleto/internal/usecase/transactions"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	envPostgresHost   = "POSTGRES_HOST"
	envPostgresPort   = "POSTGRES_PORT"
	envPostgresDbName = "POSTGRES_DB_NAME"
	envPostgresSchema = "POSTGRES_SCHEMA"
	envPostgresUser   = "POSTGRES_USER"
	envPostgresPwd    = "POSTGRES_PWD"
)

var (
	postgresHost   = os.Getenv(envPostgresHost)
	postgresPort   = os.Getenv(envPostgresPort)
	postgresDbName = os.Getenv(envPostgresDbName)
	postgresSchema = os.Getenv(envPostgresSchema)
	postgresUser   = os.Getenv(envPostgresUser)
	postgresPwd    = os.Getenv(envPostgresPwd)
)

func main() {
	databaseSettings := entity.DatabaseSettings{
		Host:   postgresHost,
		Port:   postgresPort,
		DbName: postgresDbName,
		Schema: postgresSchema,
		User:   postgresUser,
		Pwd:    postgresPwd,
	}

	databaseHelper := database.NewDatabseHelper()
	_, errConnDB := databaseHelper.InitConnection(databaseSettings)
	if errConnDB != nil {
		logrus.Fatalf("Error connecting to database: %v", errConnDB)
		return
	}

	uuidHelper := uuid.NewUuidHelper()

	accountsRepo := accounts_repository.NewAccountsRepository(databaseHelper, uuidHelper)
	transactionsRepo := transactions_repository.NewTransactionsRepository(databaseHelper, uuidHelper)
	operationsTypeRepo := operationstype.NewOperationsTypeRepository(databaseHelper)
	availableCreditLimitRepo := availablecreditlimit_repo.NewAvailableCreditLimit(databaseHelper, uuidHelper)

	accountsUseCase := accounts.NewAccountsUseCase(accountsRepo, uuidHelper)
	transactionsUseCase := transactions.NewTransactionsUseCase(transactionsRepo, accountsRepo, operationsTypeRepo, uuidHelper)
	availableCreditLimitUseCase := availablecreditlimit.NewAvailableCreditLimitUseCase(availableCreditLimitRepo, accountsRepo, uuidHelper)

	accountsHandler := handlers_accounts.NewAccountsHandler(accountsUseCase)
	transactionsHandler := handlers_transactions.NewTransactionsHandler(transactionsUseCase)
	availableCreditLimitHandler := handlers_credit.NewAvailableCreditLimit(availableCreditLimitUseCase)

	http_controller.RunHttp(accountsHandler, transactionsHandler, availableCreditLimitHandler)
}
