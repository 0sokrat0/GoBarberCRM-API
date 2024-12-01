
# ✂️ GoBarberCRM API

---

<p align="center">
  <img src="pkg/utils/img/gopher.png" alt="Gopher Logo" width="200"/>
</p>

---

![Go](https://img.shields.io/badge/Go-v1.23-blue?style=flat-square&logo=go)
![Gin](https://img.shields.io/badge/Gin-Framework-blueviolet?style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v14-blue?style=flat-square&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-green?style=flat-square&logo=swagger)

[![GitHub](https://img.shields.io/badge/GitHub-GoBarberCRM-black?style=for-the-badge&logo=github)](https://github.com/0sokrat0/GoBarberCRM-API)
[![Telegram](https://img.shields.io/badge/Telegram-sokrat_00-blue?style=for-the-badge&logo=telegram)](https://t.me/SOKRAT_00)

GoBarberCRM API — это высокопроизводительный backend API для управления системой CRM в барбершопах. Он предоставляет удобные инструменты для работы с клиентами, расписанием, бронированием услуг и уведомлениями.

---

## 📖 Возможности

- 📅 Управление расписанием сотрудников.
- 💈 Учет клиентов и их истории посещений.
- 🛠 Управление услугами и их ценами.
- 📲 Отправка уведомлений (Telegram/SMS/Email).
- 📊 Аналитика и отчеты.
- 🔒 Безопасная аутентификация через JWT.

---

## 🚀 Технологии и стек

| Технология      | Описание                                    |
|------------------|---------------------------------------------|
| **Go**          | Основной язык для разработки backend API    |
| **Gin**         | Легкий и быстрый фреймворк для REST API     |
| **PostgreSQL**  | Реляционная база данных                     |
| **GORM**        | ORM для работы с базой данных PostgreSQL    |
| **Docker**      | Контейнеризация приложения                  |
| **Swagger**     | Автодокументация API                       |

---

## 📂 Структура проекта

```plaintext
.
├── cmd
│   └── main.go
├── configs
│   ├── configs.go
│   └── config.yaml
├── docker-compose.yml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── bookings
│   │   ├── bookings_test.go
│   │   └── handler.go
│   ├── breaks
│   │   ├── breaks_test.go
│   │   └── handler.go
│   ├── clients
│   │   ├── clients_tast.go
│   │   └── handler.go
│   ├── history_log
│   ├── notifications
│   │   └── handler.go
│   ├── routes
│   │   └── routes.go
│   ├── schedules
│   │   ├── handler.go
│   │   └── schedules_test.go
│   ├── services
│   │   ├── handler.go
│   │   └── services_test.go
│   └── users
│       ├── handler.go
│       └── users_test.go
├── pkg
│   ├── db
│   │   ├── connection.go
│   │   ├── models
│   │   │   ├── bookings.go
│   │   │   ├── breaks.go
│   │   │   ├── clients.go
│   │   │   ├── history_log.go
│   │   │   ├── notifications.go
│   │   │   ├── schedules.go
│   │   │   ├── services.go
│   │   │   └── users.go
│   │   └── storage
│   │       └── storage.go
│   ├── logger
│   ├── middleware
│   │   ├── CORSM.go
│   │   └── RequestLogger.go
│   └── utils
│       ├── apiresponse.go
│       ├── hashpass.go
│       └── img
│           ├── APiGO.png
│           └── gopher.png
└── README.md

```

---

## 📂 Структура БД

<p align="center">
  <img src="pkg/utils/img/APiGO.png" alt="API Database Structure" width="600"/>
</p>

---

## 📋 Документация API

Автоматическая документация API доступна по [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

---

## 🛠 Установка и запуск

### 1. Установка зависимостей

Убедитесь, что у вас установлен Go, Docker и PostgreSQL.

```bash
go mod tidy
```

### 2. Настройка окружения

Создайте файл `config.yaml` в папке `configs/`:

```yaml
database:
  user: postgres
  password: password
  host: localhost
  port: 5432
  name: gobarbercrm

jwt:
  secret: "your_secret_key"
```

### 3. Запуск проекта

#### Локальный запуск

```bash
go run cmd/main.go
```

#### Через Docker

```bash
docker build -t gobarbercrm .
docker run -p 8080:8080 gobarbercrm
```

---

## 📂 Основные эндпоинты

| Метод   | Эндпоинт                 | Описание                                  |
|---------|--------------------------|-------------------------------------------|
| `GET`   | `/clients`              | Получить список клиентов                  |
| `POST`  | `/clients`              | Добавить нового клиента                   |
| `GET`   | `/services`             | Получить список услуг                     |
| `POST`  | `/bookings`             | Забронировать услугу                      |
| `GET`   | `/schedules`            | Получить расписание сотрудников           |

---

## 🔗 Ссылки

- 📚 [Документация API (Swagger)](http://localhost:8080/swagger/index.html)
- 📂 [GitHub Repository](https://github.com/0sokrat0/GoBarberCRM-API)

---

## 🤝 Вклад в проект

1. Форкните репозиторий.
2. Создайте ветку для своих изменений:

   ```bash
   git checkout -b feature/YourFeature
   ```

3. Сделайте коммит с изменениями:

   ```bash
   git commit -m "Add YourFeature"
   ```

4. Отправьте изменения:

   ```bash
   git push origin feature/YourFeature
   ```

5. Откройте Pull Request.

---

