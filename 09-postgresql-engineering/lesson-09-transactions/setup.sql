CREATE TABLE IF NOT EXISTS app.accounts (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    owner_name text NOT NULL,
    balance numeric(12,2) NOT NULL,
    CONSTRAINT accounts_nonnegative_balance CHECK (balance >= 0)
);

TRUNCATE app.accounts RESTART IDENTITY;

INSERT INTO app.accounts (owner_name, balance)
VALUES
    ('Sara', 500.00),
    ('Adam', 100.00);
