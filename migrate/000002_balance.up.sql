CREATE TABLE IF NOT EXISTS public.balance
(
    user_id BIGINT              NOT NULL,
    CONSTRAINT user_id PRIMARY KEY (user_id),

    balance INTEGER DEFAULT '0' NOT NULL
);