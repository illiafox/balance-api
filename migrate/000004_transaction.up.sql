CREATE TABLE IF NOT EXISTS public.transaction
(
    transaction_id BIGINT GENERATED ALWAYS AS IDENTITY,
    CONSTRAINT transaction_id PRIMARY KEY (transaction_id),
    --
    to_id          BIGINT                   NOT NULL,
    from_id        BIGINT,
    --
    action         INTEGER                  NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL
        DEFAULT (now() at time zone 'utc'),
    --
    description    TEXT                     NOT NULL
);

CREATE INDEX IF NOT EXISTS transaction_to_id_idx ON transaction (to_id);
CREATE INDEX IF NOT EXISTS transaction_from_id_idx ON transaction (from_id);