-- +goose Up
-- +goose StatementBegin

-- Типы оповещений
CREATE TYPE notification_type AS ENUM (
    'INIT_BRIEF',
    'REFRESH_BRIEF_FIRST',
    'REFRESH_BRIEF_SECOND'
);

-- Таблица шаблонов уведомлений
CREATE TABLE IF NOT EXISTS  notification_types_templates (
    notification_type notification_type NOT NULL PRIMARY KEY,
    subject_template VARCHAR(100) NOT NULL,
    body_template TEXT NOT NULL
);

-- Шаблоны для типов уведомлений
INSERT INTO notification_types_templates (notification_type, subject_template, body_template) VALUES
    ('INIT_BRIEF', 'Провести первичный инструктаж', 'Здравствуйте, {instructor_name}!\nВ подразделение "{department}" принят новый сотрудник {employee_name} на должность {position}.\nНеобходимо провести ему первичный инструктаж в течение сегодняшнего дня {today_date}.\n\nПрикрепляю лист регистрации первичного иструктажа.\n\nС уважением\nLearn-Control notification system'),
    ('REFRESH_BRIEF_FIRST', 'Провести повторный инструктаж', 'Здравствуйте, {instructor_name}!\nПриближается срок проведения повторного инструктажа на рабочем месте для сотрудника(ов) Вашего подразделения:\n{employee_name}\n\nПрикрепляю ссылку для скачивания листа регистрации повторного инструктажа.\n\n{download_link}\n\n📌 Обращаем ваше внимание: лист регистрации инструктажа необходимо скачивать в день проведения, так как при скачивании в документ автоматически подставляется текущая дата.\nКроме того, дата скачивания будет автоматически передана в систему учёта проведения инструктажей.\n\nЭто уведомление направлено за 1 месяц до наступления крайнего срока проведения повторного инструктажа. Если лист не будет скачан в течение 20 дней, повторное напоминание будет направлено за 10 дней до окончания срока.\n\nПри возникновении вопросов обращайтесь в отдел охраны труда.\n\n📩 Это письмо сформировано автоматически. Пожалуйста, не отвечайте на него.\n\nС уважением\nLearn-Control notification system'),
    ('REFRESH_BRIEF_SECOND', 'Провести повторный инструктаж', 'Здравствуйте, {instructor_name}!\nВнимание! До окончания срока проведения повторного инструктажа на рабочем месте для сотрудника(ов) Вашего подразделения осталось 10 дней:\n{employee_name}\n\nПрикрепляю ссылку для скачивания листа регистрации повторного инструктажа.\n\n{download_link}\n\n📌 Напоминаем: лист регистрации инструктажа необходимо скачивать в день проведения, так как при скачивании в документ автоматически подставляется текущая дата.\nКроме того, дата скачивания будет автоматически передана в систему учёта проведения инструктажей.\n\n‼️Просим обеспечить проведение инструктажа до {retraining_date} во избежание нарушений требований охраны труда.\n\nПри возникновении вопросов обращайтесь в отдел охраны труда.\n\n📩 Это письмо сформировано автоматически. Пожалуйста, не отвечайте на него.\n\nС уважением\nLearn-Control notification system')
ON CONFLICT DO NOTHING;

-- Таблица очереди уведомлений
CREATE TABLE IF NOT EXISTS notifications_queue (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
    training_id INTEGER NOT NULL REFERENCES trainings(id) ON DELETE CASCADE,
    notification_type notification_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    sent BOOLEAN NOT NULL DEFAULT FALSE,
    sent_at TIMESTAMP
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification_types_templates;
DROP TABLE IF EXISTS notification_queue;
DROP TYPE IF EXISTS notification_type;
-- +goose StatementEnd
