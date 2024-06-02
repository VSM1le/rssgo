-- +goose Up
create table feeds (
    id UUID PRIMARY KEY,
    create_at  TIMESTAMP NOT NULL,
    update_at  TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url text unique not null,
    user_id uuid not null references users(id) on delete cascade

);
-- +goose Down

DROP TABLE feeds;