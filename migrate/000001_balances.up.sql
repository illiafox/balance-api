CREATE TABLE IF NOT EXISTS balances
(
    user_id BIGINT UNIQUE NOT NULL,
    CONSTRAINT user_id UNIQUE(user_id),

    balance INTEGER DEFAULT '0' NOT NULL
);