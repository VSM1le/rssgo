-- +goose Up
alter table users add column api_key varchar(64) unique not NULL default (
    encode(sha256(random()::text::bytea),'hex')
);

-- +goose Down
alter table users DROP column api_key;
