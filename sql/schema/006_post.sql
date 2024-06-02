-- +goose Up
create table posts (
    id UUID PRIMARY KEY,
    create_at TIMESTAMP not null,
    update_at TIMESTAMP not null,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_id UUID not null references feeds(id) on delete cascade
);


-- +goose Down
drop table posts;