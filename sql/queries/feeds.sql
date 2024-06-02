-- name: CreateFeed :one
insert into feeds (id, create_at,update_at,name,url,user_id)
values ($1,$2,$3,$4,$5,$6)
returning *;

-- name: GetFeed :many
select * from feeds;

-- name: GetNextFeedToFetch :many
select * from feeds
order by last_fetched_at asc nulls first
limit $1;

-- name: MarkFeedAsFetch :one
update feeds
set last_fetched_at = now(),
update_at = now()
where id = $1
returning *;
