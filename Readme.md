# Learn-Control

**Learn-Control** — это веб-приложение для автоматизации процессов организации обучений и инструктажей по охране труда на предприятии. 
Система помогает отслеживать информацию о сотрудниках и их обучении, ставит задачи специалисту по охране труда и уведомляет руководителей отделов о необходимости проведения инструктажей.

## 📦 Стек технологий

**Backend:**

* Язык: Go
* Архитектура: Чистая архитектура
* gRPC (протоколы описаны в `api/learn_control.proto`)
* PostgreSQL
* Docker / Docker Compose
* Makefile для автоматизации

**Frontend:**

* React + TypeScript
* Vite
* React Router
* CSS Modules
* Взаимодействие с backend через REST API

## 🗂 Структура проекта

```
Learn-control/
├── api/                          # gRPC API
│
├── cmd/                          # Точка входа приложения
│
├── internal/                     
│   ├── adapter/                  
│   │   ├── controller/           # Контроллеры HTTP/gRPC
│   │   │   ├── files_download/   # Контроллер загрузки документов
│   │   │   │   
│   │   │   └── learn_control/    # Контроллеры бизнес-операций
│   │   │       
│   │   ├── docs_generator/       # Генерация документов
│   │   │   └── registration_form/
│   │   │       
│   │   └── notifier/             # Уведомления
│   │       └── email/
│   │
│   ├── app/                      # Инициализация зависимостей
│   │
│   ├── config/                   # Конфигурация приложения
│   │   └── modules/
│   │
│   ├── converter/                # Преобразование между слоями
│   │
│   ├── domain/                   # Доменные сущности
│   │
│   ├── logger/                   # Логирование
│   │
│   ├── service/                  # Use-case логика
│   │
│   ├── token/                    # Работа с JWT или другими токенами
│   │
│   └── worker/                   # Фоновые задачи
│       └── notification/         # Отправка уведомлений
│       │
│       └── retraining_control/   # Контроль сроков перепрохождения обучений
│
├── migrations/                   # SQL миграции
│
├── templates/                    # Шаблоны документов (регистрационные листы и др.)
├── pkg/                          # Сгенерированные файлы из proto
│   └── learn_control/
│
├── frontend/                    # Клиентская часть (React + Vite)
│   ├── public/                  # Статика
│   ├── src/
│   │   ├── app/                 # Входные компоненты
│   │   ├── pages/               # Страницы интерфейса
│   │   ├── entities/            # Типы и модели
│   │   ├── assets/              # Статика (svg, изображения)
│   │   └── main.tsx            # Вход в приложение

```

## 🚀 Быстрый старт

### 1. Клонирование репозитория

```bash
git clone https://github.com/DarYur13/Learn-control.git
cd Learn-control
```

### 2. Настройка окружения

Создайте файл `.env` по шаблону и заполните переменные 
```
# logs
LOG_LEVEL=info

# api
API_GRPC_PORT=50051
API_HTTP_PORT=8000
API_HOST=localhost

# pg
PG_HOST=localhost
PG_PORT=5432
PG_USER=postgres
PG_PASSWORD=postgres
PG_DATABASE=learn_control

# docs generator
DOCS_GENERATOR_TEMPLATE_PATH=./templates/briefings/

# notifier
EMAIL_FROM=noreply@yourdomain.com         
EMAIL_PASSWORD=yourpassword               
SMTP_HOST=smtp.yourdomain.com             
SMTP_PORT=587                             
EMAIL_USE_TLS=true                       

# notification worker
# период проверки очереди уведомлений (мин)
NOTIFICATION_WORKER_QUEUE_CHECK_PERIOD=10

# retraining control worker
# период проверки переобучений (ч)
RETRAINING_CONTROL_WORKER_QUEUE_CHECK_PERIOD=24
```

### 3. Установка зависимостей и генерация протоколов

```bash
make vendor-proto
make generate
```

### 4. Запуск проекта (Docker Compose)

```bash
docker-compose up
```

### 5. Миграции

```bash
make install-goose
make local-migration-up
```

## 📁 Полезные команды Makefile

* `make vendor-proto` — установка и настройка gRPC зависимостей
* `make generate` — генерация Go и gRPC кода из .proto
* `make install-goose` — установка утилиты для миграций
* `make local-migration-create` — создать новую миграцию
* `make local-migration-up` — применить все миграции
* `make local-migration-down` — откат миграций
* `make local-migration-status` — статус миграций
