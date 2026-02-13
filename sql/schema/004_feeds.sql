-- +goose Up
Alter table feeds add column last_fetched_at timestamp;

-- +goose Down
DROP TABLE feeds;