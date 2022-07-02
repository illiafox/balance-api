CREATE TABLE IF NOT EXISTS public.block
(
    LIKE public.balance,
    reason text                     NOT NULL,
    date   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);