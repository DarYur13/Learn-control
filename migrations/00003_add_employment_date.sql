-- +goose Up
-- +goose StatementBegin

-- Добавляем колонку с датой трудоустройства
ALTER TABLE employees
    ADD employment_date VARCHAR(10);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Удаляем колонку с датой трудоустройства
ALTER TABLE employees 
    DROP COLUMN employment_date;

-- +goose StatementEnd
