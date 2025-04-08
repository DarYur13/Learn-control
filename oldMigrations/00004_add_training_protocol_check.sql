-- +goose Up
-- +goose StatementBegin
ALTER TABLE employee_trainings
ADD has_protocol BOOLEAN;

ALTER TABLE trainings
ADD need_protocol BOOLEAN NOT NULL DEFAULT true;

CREATE OR REPLACE FUNCTION set_protocol_status()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT need_protocol FROM trainings WHERE id = NEW.training_id) THEN
        NEW.has_protocol := FALSE;
    ELSE
        NEW.has_protocol := NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_protocol_status
BEFORE INSERT ON employee_trainings
FOR EACH ROW
EXECUTE FUNCTION set_protocol_status();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trigger_set_protocol_status ON employee_trainings;
DROP FUNCTION IF EXISTS set_protocol_status();
ALTER TABLE employee_trainings
DROP COLUMN has_protocol;

ALTER TABLE trainings
DROP COLUMN need_protocol;
-- +goose StatementEnd
