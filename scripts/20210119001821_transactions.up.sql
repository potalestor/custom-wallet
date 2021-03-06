CREATE TABLE transactions(
    id bigint PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    wallet_id bigint NOT NULL REFERENCES wallets(id),
    operation integer NOT NULL CHECK(operation >= 1),
    created timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    amount numeric NOT NULL CHECK(amount > '0.00')
);

CREATE INDEX transactions_report ON transactions (wallet_id, operation, created);
