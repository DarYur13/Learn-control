-- +goose Up
-- +goose StatementBegin

-- ==============================
-- SCHEMA: EMPLOYEES & TRAININGS
-- ==============================

-- Таблица сотрудников
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    snils VARCHAR(14) UNIQUE NOT NULL,
    employment_date DATE,
    is_leader BOOLEAN NOT NULL DEFAULT FALSE
);

-- Таблица тренингов
CREATE TABLE trainings (
    id INT PRIMARY KEY,
    training VARCHAR(255) NOT NULL,
    valid_period SMALLINT,  -- NULL означает, что перепрохождение не требуется
    need_protocol BOOLEAN NOT NULL DEFAULT TRUE
);

-- Начальные данные по тренингам
INSERT INTO trainings (id, training, valid_period, need_protocol) VALUES
    (1, 'Первичный инструктаж', 6, FALSE),
    (2, 'Повторный инструктаж', 6, FALSE),
    (3, 'Вводный инструктаж', NULL, FALSE),
    (4, 'Обучение по использованию СИЗ', 36, TRUE),
    (5, 'Обучение оказанию первой помощи на производстве', 36, TRUE),
    (6, 'Обучение по общим вопросам охраны труда и функционированию системы управления охраной труда', 36, TRUE),
    (7, 'Обучение безопасным методам и приемам выполнения работ при воздействии вредных и (или) опасных производственных факторов, источников опасности, идентифицированных в рамках специальной оценки условий труда и оценки профессиональных рисков', 36, TRUE),
    (8, 'Обучение безопасным методам и приемам выполнения работ на высоте', 12, TRUE),
    (9, 'Обучение безопасным методам и приемам выполнения работ при размещении, монтаже, техническом обслуживании и ремонте технологического оборудования', 12, TRUE),
    (10, 'Обучение безопасным методам и приемам работ с ручным инструментом', 12, TRUE),
    (11, 'Обучение безопасным методам и приемам работ с радиоактивными веществами и источниками ионизирующих излучений', 12, TRUE),
    (12, 'Обучение безопасным методам и приемам работ по перемещению тяжеловесных и крупногабаритных грузов', 12, TRUE)
ON CONFLICT DO NOTHING;

-- Таблица должностей
CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    position VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    CONSTRAINT unique_position_department UNIQUE (position, department)
);

-- Связующая таблица должностей и тренингов
CREATE TABLE position_trainings (
    position_id INTEGER NOT NULL REFERENCES positions(id) ON DELETE CASCADE,
    training_id INTEGER NOT NULL REFERENCES trainings(id) ON DELETE CASCADE,
    CONSTRAINT unique_position_training UNIQUE (position_id, training_id)
);

-- Таблица прохождения обучений сотрудниками
CREATE TABLE employee_trainings (
    employee_id INT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    training_id INT NOT NULL REFERENCES trainings(id) ON DELETE CASCADE,
    training_date DATE,
    retraining_date DATE,
    has_protocol BOOLEAN,
    CONSTRAINT unique_employee_training UNIQUE (employee_id, training_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS employee_trainings;
DROP TABLE IF EXISTS position_trainings;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS trainings;
DROP TABLE IF EXISTS employees;

-- +goose StatementEnd
