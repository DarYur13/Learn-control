-- +goose Up
-- +goose StatementBegin

-- Вставка должностей
INSERT INTO positions (position, department)
VALUES
    ('Начальник отдела', 'Отдел производства средств вычислительной техники'),
    ('Кладовщик', 'Склад'),
    ('Начальник склада', 'Склад'),
    ('Мастер участка', 'Участок механической обработки'),
    ('Специалист по охране труда', 'Отдел охраны труда'),
    ('Наладчик станков с программным управлением 4 разряда', 'Участок механической обработки')
ON CONFLICT DO NOTHING;

-- Вставка сотрудников
INSERT INTO employees (full_name, birth_date, department, position, snils, employment_date, is_leader)
VALUES
    ('Иванов Алексей Петрович', '1975-03-12', 'Участок механической обработки', 'Мастер участка', '13650710080', '2025-04-07', TRUE),
    ('Смирнов Дмитрий Олегович', '1984-07-25', 'Участок механической обработки', 'Наладчик станков с программным управлением 4 разряда', '29149317528', '2025-03-19', FALSE),
    ('Попов Максим Витальевич', '1983-02-18', 'Отдел производства средств вычислительной техники', 'Начальник отдела', '88937435270', '2025-03-09', TRUE),
    ('Денисенко Дарья Юрьевна', '2001-06-13', 'Отдел охраны труда', 'Специалист по охране труда', '17503820152', '2025-03-10', FALSE),
    ('Мельников Артём Евгеньевич', '1983-12-30', 'Склад', 'Начальник склада', '12764038913', '2025-03-08', TRUE),
    ('Гаврилов Павел Владимирович', '1986-01-22', 'Склад', 'Кладовщик', '56072381734', '2025-03-27', FALSE)
ON CONFLICT DO NOTHING;

-- Заполнение position_trainings
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 1 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 4 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 5 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 6 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 10 FROM positions p WHERE p.position = 'Мастер участка' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 1 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 4 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 5 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 10 FROM positions p WHERE p.position = 'Наладчик станков с программным управлением 4 разряда' AND p.department = 'Участок механической обработки' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Начальник отдела' AND p.department = 'Отдел производства средств вычислительной техники' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 6 FROM positions p WHERE p.position = 'Начальник отдела' AND p.department = 'Отдел производства средств вычислительной техники' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Начальник отдела' AND p.department = 'Отдел производства средств вычислительной техники' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Специалист по охране труда' AND p.department = 'Отдел охраны труда' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 4 FROM positions p WHERE p.position = 'Специалист по охране труда' AND p.department = 'Отдел охраны труда' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 5 FROM positions p WHERE p.position = 'Специалист по охране труда' AND p.department = 'Отдел охраны труда' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 6 FROM positions p WHERE p.position = 'Специалист по охране труда' AND p.department = 'Отдел охраны труда' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Специалист по охране труда' AND p.department = 'Отдел охраны труда' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 4 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 5 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 6 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 12 FROM positions p WHERE p.position = 'Начальник склада' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 1 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 3 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 4 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 5 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 7 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;
INSERT INTO position_trainings (position_id, training_id)
SELECT p.id, 12 FROM positions p WHERE p.position = 'Кладовщик' AND p.department = 'Склад' ON CONFLICT DO NOTHING;


-- Обучения сотрудников
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-04-07', NULL FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 1, '2025-04-07', '2025-10-04' FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 6, '2025-04-07', '2028-03-22' FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-04-07', '2028-03-22' FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 4, '2025-04-07', '2028-03-22' FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 10, '2025-04-07', '2026-04-02' FROM employees e WHERE e.snils = '13650710080';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 5, '2025-04-07', '2028-03-22' FROM employees e WHERE e.snils = '13650710080';

INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-03-21', NULL FROM employees e WHERE e.snils = '29149317528';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 1, '2025-03-21', '2025-09-17' FROM employees e WHERE e.snils = '29149317528';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-03-21', '2028-03-05' FROM employees e WHERE e.snils = '29149317528';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 4, '2025-03-21', '2028-03-05' FROM employees e WHERE e.snils = '29149317528';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 5, '2025-03-21', '2028-03-05' FROM employees e WHERE e.snils = '29149317528';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 10, '2025-03-21', '2026-03-16' FROM employees e WHERE e.snils = '29149317528';

INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-03-15', NULL FROM employees e WHERE e.snils = '88937435270';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-03-15', '2028-02-28' FROM employees e WHERE e.snils = '88937435270';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 6, '2025-03-15', '2028-02-28' FROM employees e WHERE e.snils = '88937435270';

INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-03-10', NULL FROM employees e WHERE e.snils = '17503820152';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-03-10', '2028-02-23' FROM employees e WHERE e.snils = '17503820152';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 4, '2025-03-10', '2028-02-23' FROM employees e WHERE e.snils = '17503820152';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 5, '2025-03-10', '2028-02-23' FROM employees e WHERE e.snils = '17503820152';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 6, '2025-03-10', '2028-02-23' FROM employees e WHERE e.snils = '17503820152';

INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-03-08', NULL FROM employees e WHERE e.snils = '12764038913';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-03-08', '2028-02-20' FROM employees e WHERE e.snils = '12764038913';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 4, '2025-03-08', '2028-02-20' FROM employees e WHERE e.snils = '12764038913';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 5, '2025-03-08', '2028-02-20' FROM employees e WHERE e.snils = '12764038913';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 6, '2025-03-08', '2028-02-20' FROM employees e WHERE e.snils = '12764038913';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 12, '2025-03-08', '2026-03-03' FROM employees e WHERE e.snils = '12764038913';

INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 3, '2025-03-27', NULL FROM employees e WHERE e.snils = '56072381734';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 1, '2025-03-27', '2025-09-23' FROM employees e WHERE e.snils = '56072381734';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 7, '2025-03-27', '2028-03-11' FROM employees e WHERE e.snils = '56072381734';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 5, '2025-03-27', '2028-03-11' FROM employees e WHERE e.snils = '56072381734';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 4, '2025-03-27', '2028-03-11' FROM employees e WHERE e.snils = '56072381734';
INSERT INTO employee_trainings (employee_id, training_id, training_date, retraining_date)
SELECT e.id, 12, '2025-03-27', '2026-03-22' FROM employees e WHERE e.snils = '56072381734';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM employees WHERE snils IN (
    '13650710080',
    '29149317528',
    '88937435270',
    '17503820152',
    '12764038913',
    '56072381734'
);

DELETE FROM positions WHERE (position, department) IN (
    ('Мастер участка', 'Участок механической обработки'),
    ('Наладчик станков с программным управлением 4 разряда', 'Участок механической обработки'),
    ('Начальник отдела', 'Отдел производства средств вычислительной техники'),
    ('Специалист по охране труда', 'Отдел охраны труда'),
    ('Начальник склада', 'Склад'),
    ('Кладовщик', 'Склад')
);

-- +goose StatementEnd
