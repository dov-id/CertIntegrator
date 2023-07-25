-- +migrate Up

CREATE TYPE contracts_type_enum AS enum ('fabric', 'issuer');
CREATE TYPE tx_status_enum AS enum ('pending', 'in progress');

CREATE TABLE IF NOT EXISTS contracts (
    id      BIGSERIAL PRIMARY KEY,
    name    TEXT                  NOT NULL,
    address TEXT                  NOT NULL,
    block   BIGINT                NOT NULL,
    type    contracts_type_enum   NOT NULL
);

CREATE INDEX IF NOT EXISTS contracts_address_idx ON contracts(address);

CREATE TABLE IF NOT EXISTS users (
    address    TEXT UNIQUE NOT NULL,
    public_key TEXT        NOT NULL
);

CREATE INDEX IF NOT EXISTS users_address_idx ON users(address);

CREATE TABLE IF NOT EXISTS participants (
    user_address TEXT   NOT NULL,
    contract_id  BIGINT NOT NULL,

    UNIQUE (user_address, contract_id),

    FOREIGN KEY(user_address)
        REFERENCES users(address)
        ON DELETE CASCADE,
    FOREIGN KEY(contract_id)
        REFERENCES contracts(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS participants_user_idx ON participants(user_address);
CREATE INDEX IF NOT EXISTS participants_contract_idx ON participants(contract_id);

CREATE TABLE IF NOT EXISTS transactions (
    id       BIGSERIAL      PRIMARY KEY,
    status   tx_status_enum              NOT NULL,
    course   TEXT                        NOT NULL,
    state    BYTEA                       NOT NULL
);

CREATE INDEX IF NOT EXISTS transactions_type_idx ON transactions(status);

CREATE TABLE IF NOT EXISTS tx_statuses (
    tx_id   BIGINT NOT NULL,
    network TEXT   NOT NULL,

    FOREIGN KEY(tx_id)
        REFERENCES transactions(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS tx_statuses_tx_id_idx ON tx_statuses(tx_id);

CREATE TABLE IF NOT EXISTS mt_nodes (
    mt_id      BIGINT,
    key        BYTEA,
    type       SMALLINT NOT NULL,
    child_l    BYTEA,
    child_r    BYTEA,
    entry      BYTEA,
    created_at BIGINT,
    deleted_at BIGINT,

    PRIMARY KEY(mt_id, key)
);

CREATE TABLE IF NOT EXISTS mt_roots (
    mt_id      BIGINT PRIMARY KEY,
    key        BYTEA,
    created_at BIGINT,
    deleted_at BIGINT
);

-- +migrate Down

DROP TABLE IF EXISTS mt_roots;
DROP TABLE IF EXISTS mt_nodes;

DROP TABLE IF EXISTS tx_statuses;
DROP INDEX IF EXISTS tx_statuses_tx_id_idx;

DROP INDEX IF EXISTS transactions_type_idx;
DROP TABLE IF EXISTS transactions;

DROP INDEX IF EXISTS participants_user_idx;
DROP INDEX IF EXISTS participants_contract_idx;
DROP TABLE IF EXISTS participants;

DROP INDEX IF EXISTS users_address_idx;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS contracts_address_idx;
DROP TABLE IF EXISTS contracts;

DROP TYPE IF EXISTS tx_status_enum;
DROP TYPE IF EXISTS contracts_type_enum;