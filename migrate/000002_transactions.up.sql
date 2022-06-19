CREATE TABLE IF NOT EXISTS transactions
(
    -- table

    transaction_id BIGINT GENERATED ALWAYS AS IDENTITY,
    to_id          BIGINT                  NOT NULL,
    from_id        BIGINT,
    action         INTEGER                 NOT NULL,
    created_at           TIMESTAMP WITH TIME ZONE
        DEFAULT (now() at time zone 'utc') NOT NULL,
    description    TEXT                    NOT NULL,

    -- constraints

    CONSTRAINT fk_balances
        FOREIGN KEY (to_id)
            REFERENCES balances (balance_id)
            ON UPDATE RESTRICT
            ON DELETE RESTRICT,
        FOREIGN KEY (from_id)
            REFERENCES balances (balance_id)
            ON UPDATE RESTRICT
            ON DELETE RESTRICT
);