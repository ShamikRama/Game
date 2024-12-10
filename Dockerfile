# Используем базовый образ Go
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Устанавливаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o myapp .

# Устанавливаем golang-migrate
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate

# Указываем переменную окружения для пути к конфигу
ENV CONFIG_PATH=/app/config/config.yaml

# Команда для запуска приложения
CMD ["./myapp"]