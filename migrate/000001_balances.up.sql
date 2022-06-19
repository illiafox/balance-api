CREATE TABLE IF NOT EXISTS balances
(
    -- table

    balance_id BIGINT UNIQUE GENERATED ALWAYS AS IDENTITY,
    user_id    BIGINT UNIQUE       NOT NULL,
    balance    INTEGER DEFAULT '0' NOT NULL
);