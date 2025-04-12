-- +goose Up
-- +goose StatementBegin

-- Enum тип задач
CREATE TYPE task_type AS ENUM (
    'PROVIDE', 
    'ASSIGN', 
    'CHOOSE', 
    'SET', 
    'CONFIRM', 
    'CONTROL'
);

-- Таблица текстов типов задач
CREATE TABLE IF NOT EXISTS task_types_texts (
    task_type task_type NOT NULL PRIMARY KEY,
    task_text VARCHAR(255) NOT NULL
);

-- Таблица задач
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    task_type task_type NOT NULL REFERENCES task_types_texts(task_type),
    training_id INTEGER REFERENCES trainings(id),
    employee_id BIGINT REFERENCES employees(id),
    executor_id BIGINT REFERENCES employees(id),
    position_id BIGINT REFERENCES positions(id),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    done BOOLEAN NOT NULL DEFAULT false,
    done_at DATE
);

-- Ограничение: одна задача типа CHOOSE на должность
CREATE UNIQUE INDEX IF NOT EXISTS unique_choose_task_per_position
    ON tasks(position_id)
    WHERE task_type = 'CHOOSE';

-- Начальные данные типов задач
INSERT INTO task_types_texts (task_type, task_text) VALUES
    ('PROVIDE', 'Провести обучение'),
    ('ASSIGN', 'Назначить обучение'),
    ('CHOOSE', 'Определить перечень необходимых обучающих мероприятий для новой должности'),
    ('SET', 'Внести дату проверки знаний'),
    ('CONFIRM', 'Подтвердить получение протокола'),
    ('CONTROL', 'Проконтролировать проведение обучения')
ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS unique_choose_task_per_position;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS task_types_texts;
DROP TYPE IF EXISTS task_type;

-- +goose StatementEnd
