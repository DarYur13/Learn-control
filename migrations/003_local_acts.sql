-- +goose Up
-- +goose StatementBegin

-- ==================
-- SCHEMA: LOCAL ACTS
-- ==================

CREATE TABLE local_acts (
    id SERIAL PRIMARY KEY,
    act_name VARCHAR(255) NOT NULL
);

-- Начальные данные
INSERT INTO local_acts (act_name) VALUES
    ('Положение об организации обучения по охране труда работников АО «N»')
ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS local_acts;

-- +goose StatementEnd

