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

# Указываем переменную окружения для пути к конфигу
ENV CONFIG_PATH=/app/config/config.yaml

# Команда для запуска приложения
CMD ["./myapp"]