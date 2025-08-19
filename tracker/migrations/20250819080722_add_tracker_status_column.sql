-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE tracker ADD COLUMN status VARCHAR(50) NOT NULL DEFAULT 'active';
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
ALTER TABLE tracker DROP COLUMN status;
