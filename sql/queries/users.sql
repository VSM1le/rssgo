-- name: CreateUser :one
insert into users (id, create_at,update_at,name,api_key)
values ($1,$2,$3,$4,
        encode(sha256(random()::text::bytea),'hex')
)
returning *;

-- name: GetUserByAPIKey :one
select * from users where api_key = $1;