-- +goose Up
-- +goose StatementBegin

CREATE TYPE task_type AS ENUM ('PROVIDE', 'ASSIGN', 'CHOOSE', 'SET', 'CONFIRM');

-- Создание таблицы типов задач
CREATE TABLE IF NOT EXISTS task_types_texts (
    task_type task_type NOT NULL PRIMARY KEY,
    template_text VARCHAR(255) NOT NULL
);

-- Создание таблицы задач
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    training_id INTEGER REFERENCES trainings(id),
    employee_id BIGINT REFERENCES employees(id),
    executor_id BIGINT REFERENCES employees(id),
    position_id BIGINT REFERENCES positions(id),
    task_type task_type NOT NULL REFERENCES task_types_texts(task_type),
    done BOOLEAN NOT NULL DEFAULT false,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    done_at DATE
);

INSERT INTO task_types_texts (task_type, template_text) VALUES
    ('PROVIDE', 'провести @training'),
    ('ASSIGN', 'назначить @training'),
    ('CHOOSE', 'определить перечень необходимых обучающих мероприятий для новой должности'),
    ('SET', 'внести дату, когда сотрудник прошел @training'),
    ('CONFIRM', 'подтвердить получение протокола')
ON CONFLICT DO NOTHING;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Удаление таблицs задач
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS task_types_texts;
DROP TYPE IF EXISTS task_type;

-- +goose StatementEnd
