CREATE SCHEMA IF NOT EXISTS "pismo";
GRANT ALL ON SCHEMA pismo to postgres;

CREATE TABLE IF NOT EXISTS pismo.accounts (
    "account_id" uuid NOT NULL,
    "document_number" varchar(255) NOT NULL,
    PRIMARY KEY ("account_id")
);

CREATE INDEX ON pismo.accounts ("account_id");

INSERT INTO pismo.accounts ("account_id", "document_number") VALUES ('d2ba4e0d-b834-4c14-9bd6-3488fc619f5a', '12345678');

CREATE TABLE IF NOT EXISTS pismo.operationsTypes (
    "operation_type_id" SERIAL,
    "description" varchar(255) NOT NULL,
    PRIMARY KEY ("operation_type_id")
);

CREATE INDEX ON pismo.operationsTypes ("operation_type_id");

INSERT INTO pismo.operationsTypes ("description") VALUES ('CASH PURCHASE');
INSERT INTO pismo.operationsTypes ("description") VALUES ('INSTALLMENT PURCHASE');
INSERT INTO pismo.operationsTypes ("description") VALUES ('WITHDRAWAL');
INSERT INTO pismo.operationsTypes ("description") VALUES ('PAYMENT');

CREATE TABLE IF NOT EXISTS pismo.transactions (
    "transaction_id" uuid NOT NULL,
    "account_id" uuid NOT NULL,
    "operation_type_id" int NOT NULL,
    "amount" numeric(10,2) NOT NULL,
    "event_date" timestamp NOT NULL,
    PRIMARY KEY ("transaction_id")
);

CREATE INDEX ON pismo.transactions ("transaction_id");
CREATE INDEX ON pismo.transactions ("account_id");
CREATE INDEX ON pismo.transactions ("operation_type_id");