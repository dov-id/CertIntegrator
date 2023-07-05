-- +migrate Up

create type contracts_type_enum as enum ('fabric', 'issuer');

create table if not exists contracts (
    id bigserial primary key,
    name text not null,
    address text not null,
    block bigint not null,
    type contracts_type_enum not null,
    unique(address)
);

create index contracts_address_idx on contracts(address);

create table if not exists mt_nodes (
    mt_id bigint,
    key bytea,
    type smallint not null,
    child_l bytea,
    child_r bytea,
    entry bytea,
    created_at bigint,
    deleted_at bigint,
    primary key(mt_id, key)
);

create table if not exists mt_roots (
    mt_id bigint primary key,
    key bytea,
    created_at bigint,
    deleted_at bigint
);

-- +migrate Down

drop table if exists mt_roots;
drop table if exists mt_nodes;
drop index if exists contracts_address_idx;
drop table if exists contracts;
drop type if exists contracts_type_enum;