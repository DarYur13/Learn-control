-- +goose Up
-- +goose StatementBegin

-- Создание таблицы локальных актов
CREATE TABLE IF NOT EXISTS local_acts (
    id SERIAL PRIMARY KEY,
    act_name VARCHAR(255) NOT NULL
);

INSERT INTO local_acts (act_name) VALUES(
    ('Положение об организации обучения по охране труда работников АО «N»')
)
ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE local_acts
-- +goose StatementEnd
