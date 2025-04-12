-- +goose Up
-- +goose StatementBegin

-- Таблица актов
CREATE TABLE IF NOT EXISTS local_acts (
    id SERIAL PRIMARY KEY,
    act_name VARCHAR(255) NOT NULL
);

-- Начальные данные
INSERT INTO local_acts (act_name) VALUES
    ('Положение об организации обучения по охране труда работников АО «N»')
ON CONFLICT DO NOTHING;

-- Таблица соотвествия актов обучениям
CREATE TABLE IF NOT EXISTS  acts_trainings (
    local_act_id     INT NOT NULL REFERENCES local_acts(id),
    training_id      INT NOT NULL REFERENCES trainings(id)
);

-- Начальные данные для таблицы соотвествия
INSERT INTO acts_trainings (local_act_id, training_id) VALUES
    (1, 1),
    (1, 2)
ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS acts_trainings;
DROP TABLE IF EXISTS local_acts;

-- +goose StatementEnd

