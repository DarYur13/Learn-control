-- +goose Up
-- +goose StatementBegin
ALTER TABLE employees
ADD is_leader BOOLEAN;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE employees 
DROP COLUMN is_leader;
-- +goose StatementEnd
