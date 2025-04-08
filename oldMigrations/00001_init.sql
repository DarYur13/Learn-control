-- +goose Up
-- +goose StatementBegin

-- Создание таблицы сотрудников
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    birth_date VARCHAR(10) NOT NULL,
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    snils VARCHAR(14) UNIQUE NOT NULL
);

-- Создание таблицы тренингов
CREATE TABLE IF NOT EXISTS trainings (
    id SERIAL PRIMARY KEY,
    training VARCHAR(255) NOT NULL,
    valid_period SMALLINT  -- NULL означает, что перепрохождение не требуется
);

-- Вставка данных о тренингах
INSERT INTO trainings (training, valid_period) VALUES
    ('первичный инструктаж', 6),
    ('вводный инструктаж', NULL),
    ('обучение по использованию СИЗ', 36),
    ('обучение оказанию первой помощи на производстве', 36),
    ('обучение по общим вопросам охраны труда и функционированию системы управления охраной труда', 36),
    ('обучение безопасным методам и приемам выполнения работ при воздействии вредных и (или) опасных производственных факторов, источников опасности, идентифицированных в рамках специальной оценки условий труда и оценки профессиональных рисков', 36),
    ('обучение безопасным методам и приемам выполнения работ повышенной опасности', 12)
ON CONFLICT DO NOTHING;

-- Создание таблицы должностей
CREATE TABLE IF NOT EXISTS positions (
    id SERIAL PRIMARY KEY,
    position VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    CONSTRAINT unique_position_department UNIQUE (position, department)
);

-- Связующая таблица должностей и тренингов
CREATE TABLE IF NOT EXISTS position_trainings (
    position_id INTEGER NOT NULL,
    training_id INTEGER NOT NULL,
    CONSTRAINT fk_position FOREIGN KEY (position_id) REFERENCES positions(id) ON DELETE CASCADE,
    CONSTRAINT fk_training FOREIGN KEY (training_id) REFERENCES trainings(id) ON DELETE CASCADE,
    CONSTRAINT unique_position_training UNIQUE (position_id, training_id)
);

-- Таблица прохождения обучений сотрудниками
CREATE TABLE IF NOT EXISTS employee_trainings (
    employee_id INT NOT NULL,
    training_id INT NOT NULL,
    training_date DATE NOT NULL DEFAULT CURRENT_DATE,
    retraining_date DATE,
    CONSTRAINT fk_employee FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    CONSTRAINT fk_training FOREIGN KEY (training_id) REFERENCES trainings(id) ON DELETE CASCADE,
    CONSTRAINT unique_employee_training UNIQUE (employee_id, training_id)
);

-- Функция для установки retraining_date на основе valid_period
CREATE OR REPLACE FUNCTION set_retraining_date() RETURNS TRIGGER AS $$
DECLARE
    period SMALLINT;
BEGIN
    SELECT valid_period INTO period FROM trainings WHERE id = NEW.training_id;

    IF period IS NOT NULL THEN
        NEW.retraining_date := NEW.training_date + (period || ' months')::INTERVAL;
    ELSE
        NEW.retraining_date := NULL;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для установки retraining_date перед вставкой
CREATE TRIGGER trg_set_retraining_date
BEFORE INSERT ON employee_trainings
FOR EACH ROW EXECUTE FUNCTION set_retraining_date();

-- Функция для автозаполнения employee_trainings при добавлении сотрудника
CREATE OR REPLACE FUNCTION add_employee_trainings() RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO employee_trainings (employee_id, training_id)
    SELECT NEW.id, pt.training_id
    FROM positions p
    JOIN position_trainings pt ON p.id = pt.position_id
    WHERE p.position = NEW.position AND p.department = NEW.department;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для добавления обучений при вставке нового сотрудника
CREATE TRIGGER trg_add_employee_trainings
AFTER INSERT ON employees
FOR EACH ROW EXECUTE FUNCTION add_employee_trainings();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаление триггеров и функций
DROP TRIGGER IF EXISTS trg_set_retraining_date ON employee_trainings;
DROP FUNCTION IF EXISTS set_retraining_date();

DROP TRIGGER IF EXISTS trg_add_employee_trainings ON employees;
DROP FUNCTION IF EXISTS add_employee_trainings();

-- Удаление данных о тренингах
DELETE FROM trainings WHERE training IN (
    'первичный инструктаж',
    'вводный инструктаж',
    'обучение по использованию СИЗ',
    'обучение оказанию первой помощи на производстве',
    'обучение по общим вопросам охраны труда и функционированию системы управления охраной труда',
    'обучение безопасным методам и приемам выполнения работ при воздействии вредных и (или) опасных производственных факторов, источников опасности, идентифицированных в рамках специальной оценки условий труда и оценки профессиональных рисков',
    'обучение безопасным методам и приемам выполнения работ повышенной опасности'
);

-- Удаление таблиц
DROP TABLE IF EXISTS employee_trainings;
DROP TABLE IF EXISTS position_trainings;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS trainings;
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
