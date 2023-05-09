-- +migrate Up

create table if not exists contract_addresses (
    address text not null,
    unique(address)
);

create index contract_addresses_address_ids on contract_addresses(address);

create table if not exists blocks (
    contract_name text not null,
    last_block_number bigint not null,
    unique(contract_name)
);

-- +migrate Down

drop table if exists blocks;
drop index if exists contract_addresses_address_ids;
drop table if exists contract_addresses;