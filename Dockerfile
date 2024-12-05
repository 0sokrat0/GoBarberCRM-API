# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Установка необходимых пакетов
RUN apk update && apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование go.mod и go.sum
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование всего исходного кода
COPY . .

# Проверка наличия конфигурационного файла
RUN ls -la /app/app/configs/

# Сборка приложения
RUN go build -o gobarbercrm app/cmd/main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Установка необходимых пакетов
RUN apk --no-cache add ca-certificates

# Установка рабочей директории
WORKDIR /root/app

# Копирование бинарника из билдера
COPY --from=builder /app/gobarbercrm .

# Создание директории для конфигурации
RUN mkdir -p ./configs

# Копирование конфигурационного файла с правильным путём
COPY --from=builder /app/app/configs/config.yaml ./configs/

# Проверка наличия конфигурационного файла в runtime-стадии
RUN ls -la ./configs/

# Открытие порта приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["./gobarbercrm"]
