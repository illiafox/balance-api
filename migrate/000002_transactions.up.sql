CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id BIGINT GENERATED ALWAYS AS IDENTITY,
    --
    to_id          BIGINT                   NOT NULL,
    CONSTRAINT to_id FOREIGN KEY (to_id) REFERENCES balances (user_id)
        ON UPDATE RESTRICT ON DELETE RESTRICT,
    --
    from_id        BIGINT,
    CONSTRAINT from_id FOREIGN KEY (from_id) REFERENCES balances (user_id)
        ON UPDATE RESTRICT ON DELETE RESTRICT,
    --
    action         INTEGER                  NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL
        DEFAULT (now() at time zone 'utc'),
    description    TEXT                     NOT NULL
);