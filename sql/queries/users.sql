-- name: CreateUser :one


INSERT INTO users (id, created_at, updated_at, name, api_key)
--VALUES ($1, $2, $3, $4, $5)  -- We could do like this, but instead we want SQL to do this. We do not want to treat api_key like others parameters
VALUES ($1, $2, $3, $4,
    encode(sha256(random()::text::bytea), 'hex')
    )
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;

