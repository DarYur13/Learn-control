-- +goose Up
-- +goose StatementBegin

-- Таблица сотрудников
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    snils VARCHAR(14) UNIQUE NOT NULL,
    employment_date DATE NOT NULL,
    is_leader BOOLEAN NOT NULL DEFAULT FALSE,
    email VARCHAR(100) UNIQUE
);

-- Типы тренингов
CREATE TYPE training_type AS ENUM (
    'INTRODUCTORY',
    'INITIAL',
    'REFRESHER',
    'REGULAR'
);

-- Таблица тренингов
CREATE TABLE IF NOT EXISTS trainings (
    id SERIAL PRIMARY KEY,
    training_type training_type NOT NULL,
    training_name VARCHAR(255) NOT NULL,
    valid_period SMALLINT  -- NULL означает, что перепрохождение не требуется
);

-- Начальные данные по обучениям
INSERT INTO trainings (training_type, training_name, valid_period) VALUES
    ('INITIAL', 'Первичный инструктаж', 6),
    ('REFRESHER', 'Повторный инструктаж', 6),
    ('INTRODUCTORY', 'Вводный инструктаж', NULL),
    ('REGULAR', 'Обучение по использованию СИЗ', 36),
    ('REGULAR', 'Обучение оказанию первой помощи на производстве', 36),
    ('REGULAR', 'Обучение по общим вопросам охраны труда и функционированию системы управления охраной труда', 36),
    ('REGULAR', 'Обучение безопасным методам и приемам выполнения работ при воздействии вредных и (или) опасных производственных факторов, источников опасности, идентифицированных в рамках специальной оценки условий труда и оценки профессиональных рисков', 36),
    ('REGULAR', 'Обучение безопасным методам и приемам выполнения работ на высоте', 12),
    ('REGULAR', 'Обучение безопасным методам и приемам выполнения работ при размещении, монтаже, техническом обслуживании и ремонте технологического оборудования', 12),
    ('REGULAR', 'Обучение безопасным методам и приемам работ с ручным инструментом', 12),
    ('REGULAR', 'Обучение безопасным методам и приемам работ с радиоактивными веществами и источниками ионизирующих излучений', 12),
    ('REGULAR', 'Обучение безопасным методам и приемам работ по перемещению тяжеловесных и крупногабаритных грузов', 12)
ON CONFLICT DO NOTHING;

-- Таблица должностей
CREATE TABLE IF NOT EXISTS positions (
    id SERIAL PRIMARY KEY,
    position VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    CONSTRAINT unique_position_department UNIQUE (position, department)
);

-- Связующая таблица должностей и обучений
CREATE TABLE IF NOT EXISTS position_trainings (
    position_id INTEGER NOT NULL REFERENCES positions(id) ON DELETE CASCADE,
    training_id INTEGER NOT NULL REFERENCES trainings(id) ON DELETE CASCADE,
    CONSTRAINT unique_position_training UNIQUE (position_id, training_id)
);

-- Таблица прохождения обучений сотрудниками
CREATE TABLE IF NOT EXISTS employee_trainings (
    employee_id BIGINT NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    training_id INT NOT NULL REFERENCES trainings(id) ON DELETE CASCADE,
    training_date DATE,
    retraining_date DATE,
    has_protocol BOOLEAN DEFAULT NULL,
    CONSTRAINT unique_employee_training UNIQUE (employee_id, training_id, training_date)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS employee_trainings;
DROP TABLE IF EXISTS position_trainings;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS trainings;
DROP TABLE IF EXISTS employees;
DROP TYPE IF EXISTS training_type;


-- +goose StatementEnd
