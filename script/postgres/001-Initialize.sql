CREATE SCHEMA IF NOT EXISTS "pismo";
GRANT ALL ON SCHEMA pismo to postgres;

CREATE TYPE OperationsTypes AS ENUM ('CASH PURCHASE', 'INSTALLMENT PURCHASE', 'WITHDRAWAL', 'PAYMENT');

CREATE TABLE IF NOT EXISTS pismo.accounts (
    "account_ID" uuid NOT NULL,
    "document_number" varchar(255) NOT NULL,
    PRIMARY KEY ("account_ID")
);

CREATE INDEX ON pismo.accounts ("account_ID");

CREATE TABLE IF NOT EXISTS pismo.operationsTypes (
    "operations_type_ID" SERIAL,
    "document_number" varchar(255) NOT NULL,
    PRIMARY KEY ("operations_type_ID")
);

CREATE INDEX ON pismo.operationsTypes ("operations_type_ID");

INSERT INTO pismo.operationsTypes ("document_number") VALUES ('CASH PURCHASE');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('INSTALLMENT PURCHASE');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('WITHDRAWAL');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('PAYMENT');

CREATE TABLE IF NOT EXISTS pismo.transactions (
    "transaction_ID" uuid NOT NULL,
    "account_ID" uuid NOT NULL,
    "operations_type_ID" int NOT NULL,
    "amount" numeric(10,2) NOT NULL,
    "event_date" timestamp NOT NULL,
    PRIMARY KEY ("transaction_ID")
);

CREATE INDEX ON pismo.transactions ("transaction_ID");
CREATE INDEX ON pismo.transactions ("account_ID");
CREATE INDEX ON pismo.transactions ("operations_type_ID");