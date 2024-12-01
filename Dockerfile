# Используем базовый образ Go
FROM golang:1.23.3-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка приложения
RUN go build -o main cmd/main.go

# Экспонируем порт
EXPOSE 8080

# Запуск приложения
CMD ["./main"]
