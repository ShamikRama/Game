# ВНИМАНИЕ - собрать файл находясь в корне проекта с помощью docker build -t myapp
# Используем базовый образ Go
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем только go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем содержимое папки cmd/ в рабочую директорию
COPY cmd/ ./cmd/

# Копируем остальные необходимые файлы (например, internal/)
COPY internal/ ./internal/

# Копируем конфигурационные файлы
COPY config/ ./config/

# Копируем содержимое папки pkg/ в рабочую директорию
COPY pkg/ ./pkg/

# Собираем приложение
RUN go build -o myapp ./cmd/

# Устанавливаем golang-migrate
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

# Указываем переменную окружения для пути к конфигу
ENV CONFIG_PATH=/app/config/local.yaml

# Команда для запуска приложения
CMD ["./myapp"]