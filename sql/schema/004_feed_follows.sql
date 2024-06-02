-- +goose Up
create table feed_follows (
    id UUID PRIMARY KEY,
    create_at  TIMESTAMP NOT NULL,
    update_at  TIMESTAMP NOT NULL,
    user_id UUID not null references users(id) on delete cascade,
    feed_id UUID NOT null references feeds(id) on delete cascade,
    unique(user_id,feed_id)

);
-- +goose Down

DROP TABLE feed_follows;