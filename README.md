
# ✂️ GoBarberCRM API

![Go](https://img.shields.io/badge/Go-v1.20-blue?style=flat-square&logo=go)
![Gin](https://img.shields.io/badge/Gin-Framework-blueviolet?style=flat-square)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v14-blue?style=flat-square&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-API%20Docs-green?style=flat-square&logo=swagger)

[![GitHub](https://img.shields.io/badge/GitHub-GoBarberCRM-black?style=for-the-badge&logo=github)](https://github.com/0sokrat0/GoBarberCRM-API)
[![Telegram](https://img.shields.io/badge/Telegram-Community-blue?style=for-the-badge&logo=telegram)](https://t.me/SOKRAT_00)

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
├── cmd/
│   └── main.go                 # Главный файл запуска API
├── configs/
│   └── configs.go              # Настройки приложения
├── docs/                       # Документация API (Swagger, OpenAPI)
├── internal/                   # Основная логика приложения
│   ├── appointments/           # Работа с бронированиями
│   ├── clienthistory/          # История клиентов
│   ├── clients/                # Управление клиентами
│   ├── department/             # Филиалы/отделы
│   ├── notifications/          # Уведомления
│   ├── schedules/              # Расписания сотрудников
│   ├── services/               # Управление услугами
│   └── users/                  # Управление пользователями
└── pkg/                        # Общие пакеты
    ├── db/                     # Подключение к базе данных
    ├── logger/                 # Логирование
    └── utils/                  # Вспомогательные функции
```

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

#### Локальный запуск:
```bash
go run cmd/main.go
```

#### Через Docker:
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
- 📂 [GitHub Repository](https://github.com/your-repo/gobarbercrm)
- 🐋 [Docker Hub](https://hub.docker.com/repository/docker/your-repo/gobarbercrm)

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
