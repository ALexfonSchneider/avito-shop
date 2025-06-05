-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id            TEXT PRIMARY KEY,
    username      TEXT UNIQUE NOT NULL,
    password_hash TEXT        NOT NULL,
    balance       INTEGER DEFAULT 0 CONSTRAINT balance_more_then_zero CHECK ( balance >= 0 ),
    created_at    TIMESTAMP
);

CREATE TABLE transactions
(
    id          TEXT PRIMARY KEY,
    sender_id   TEXT REFERENCES users (id),
    receiver_id TEXT REFERENCES users (id),
    amount      INTEGER NOT NULL,
    created_at  TIMESTAMP
);

CREATE INDEX idx_transactions_sender ON transactions (sender_id);
CREATE INDEX idx_transactions_receiver ON transactions (receiver_id);

CREATE TABLE merch
(
    id          TEXT PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    price       INTEGER      NOT NULL CONSTRAINT price_is_positive CHECK ( price > 0 ),
    created_at  TIMESTAMP
);

CREATE INDEX idx_merch_name ON merch (name);

CREATE TABLE purchases
(
    id           TEXT PRIMARY KEY,
    user_id      TEXT REFERENCES users (id),
    merch_id     TEXT REFERENCES merch (id),
    amount      INTEGER NOT NULL CONSTRAINT amount_is_positive CHECK ( amount >= 0 ),
    quantity     INTEGER NOT NULL CONSTRAINT quantity_is_positive CHECK ( quantity >= 0 ),
    purchased_at TIMESTAMP
);

CREATE INDEX idx_purchases_user ON purchases (user_id);
CREATE INDEX idx_purchases_merch ON purchases (merch_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE purchases;
DROP TABLE merch;
DROP TABLE transactions;
DROP TABLE users;
-- +goose StatementEnd
