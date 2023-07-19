-- +migrate Up

CREATE TYPE contracts_type_enum AS enum ('fabric', 'issuer');

CREATE TABLE IF NOT EXISTS contracts (
    id      BIGSERIAL PRIMARY KEY,
    name    TEXT                  NOT NULL,
    address TEXT                  NOT NULL,
    block   BIGINT                NOT NULL,
    type    contracts_type_enum   NOT NULL
);

CREATE INDEX contracts_address_idx ON contracts(address);

CREATE TABLE IF NOT EXISTS users (
    address    TEXT UNIQUE NOT NULL,
    public_key TEXT        NOT NULL
);

CREATE INDEX users_address_idx ON users(address);

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
DROP INDEX IF EXISTS users_address_idx;
DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS contracts_address_idx;
DROP TABLE IF EXISTS contracts;
DROP TYPE IF EXISTS contracts_type_enum;