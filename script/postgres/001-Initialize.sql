CREATE SCHEMA IF NOT EXISTS "pismo";
GRANT ALL ON SCHEMA pismo to postgres;

CREATE TABLE IF NOT EXISTS pismo.accounts (
    "account_id" uuid NOT NULL,
    "document_number" varchar(255) NOT NULL,
    PRIMARY KEY ("account_id")
);

CREATE INDEX ON pismo.accounts ("account_id");

CREATE TABLE IF NOT EXISTS pismo.operationsTypes (
    "operations_type_id" SERIAL,
    "document_number" varchar(255) NOT NULL,
    PRIMARY KEY ("operations_type_id")
);

CREATE INDEX ON pismo.operationsTypes ("operations_type_id");

INSERT INTO pismo.operationsTypes ("document_number") VALUES ('CASH PURCHASE');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('INSTALLMENT PURCHASE');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('WITHDRAWAL');
INSERT INTO pismo.operationsTypes ("document_number") VALUES ('PAYMENT');

CREATE TABLE IF NOT EXISTS pismo.transactions (
    "transaction_id" uuid NOT NULL,
    "account_id" uuid NOT NULL,
    "operations_type_id" int NOT NULL,
    "amount" numeric(10,2) NOT NULL,
    "event_date" timestamp NOT NULL,
    PRIMARY KEY ("transaction_id")
);

CREATE INDEX ON pismo.transactions ("transaction_id");
CREATE INDEX ON pismo.transactions ("account_id");
CREATE INDEX ON pismo.transactions ("operations_type_id");