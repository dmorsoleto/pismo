CREATE TABLE IF NOT EXISTS pismo.availableCreditLimit (
    "available_credit_limit_id" uuid NOT NULL,
    "account_id" uuid NOT NULL,
    "available_credit_limit" numeric(10, 2) NOT NULL,
    PRIMARY KEY ("available_credit_limit_id")
);

CREATE INDEX ON pismo.availableCreditLimit ("available_credit_limit_id");
CREATE INDEX ON pismo.availableCreditLimit ("account_id");