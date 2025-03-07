-- +goose Up
-- +goose StatementBegin

-- Удаляем значение по умолчанию и делаем training_date NULLABLE
ALTER TABLE employee_trainings 
    ALTER COLUMN training_date DROP DEFAULT,
    ALTER COLUMN training_date DROP NOT NULL;

-- Обновляем функцию расчета retraining_date, чтобы она срабатывала при изменении training_date
CREATE OR REPLACE FUNCTION update_retraining_date() RETURNS TRIGGER AS $$
DECLARE
    period SMALLINT;
BEGIN
    IF NEW.training_date IS NOT NULL THEN
        SELECT valid_period INTO period FROM trainings WHERE id = NEW.training_id;

        IF period IS NOT NULL THEN
            NEW.retraining_date := NEW.training_date + (period || ' months')::INTERVAL;
        ELSE
            NEW.retraining_date := NULL;
        END IF;
    ELSE
        NEW.retraining_date := NULL;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Пересоздаем триггер, который срабатывает при вставке и обновлении training_date
DROP TRIGGER IF EXISTS trg_set_retraining_date ON employee_trainings;
CREATE TRIGGER trg_set_retraining_date
BEFORE INSERT OR UPDATE OF training_date ON employee_trainings
FOR EACH ROW EXECUTE FUNCTION update_retraining_date();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Восстанавливаем training_date как NOT NULL и возвращаем значение по умолчанию
ALTER TABLE employee_trainings 
    ALTER COLUMN training_date SET NOT NULL,
    ALTER COLUMN training_date SET DEFAULT CURRENT_DATE;

-- Удаляем триггер и функцию
DROP TRIGGER IF EXISTS trg_set_retraining_date ON employee_trainings;
DROP FUNCTION IF EXISTS update_retraining_date();

-- +goose StatementEnd
