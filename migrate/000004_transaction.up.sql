CREATE TABLE IF NOT EXISTS public.transaction
(
    transaction_id BIGINT GENERATED ALWAYS AS IDENTITY,
    --
    to_id          BIGINT                   NOT NULL,
 --   CONSTRAINT to_id FOREIGN KEY (to_id) REFERENCES balance (user_id)
 --       ON UPDATE NO ACTION ON DELETE NO ACTION,
    --
    from_id        BIGINT,
 --   CONSTRAINT from_id FOREIGN KEY (from_id) REFERENCES balance (user_id)
 --       ON UPDATE NO ACTION ON DELETE NO ACTION,
    --
    action         INTEGER                  NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE NOT NULL
        DEFAULT (now() at time zone 'utc'),
    description    TEXT                     NOT NULL
);