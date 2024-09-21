-- +goose Up

ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex') -- We need that so our existing users will have this unique not null api key 
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;