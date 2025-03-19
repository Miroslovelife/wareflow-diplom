# Используем официальный образ Go как базовый
FROM golang:1.23.4-alpine AS builder

RUN apk --no-cache add make git

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /whareflow-server

# Копируем go.mod и go.sum (если они есть), чтобы Docker кэшировал зависимости
COPY go.mod go.sum ./

# Загружаем зависимости (это кеширует зависимости, чтобы не тянуть их каждый раз)
RUN go mod tidy

# Копируем весь исходный код в контейнер
COPY . .

# Собираем приложение
RUN make build

# Используем минимальный образ для финального контейнера
FROM alpine:latest

# Устанавливаем необходимые библиотеки для работы Go-приложения (например, libc)
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из контейнера сборки
COPY --from=builder /whareflow-server/cmd/wareflow /usr/local/bin/wareflow

# Устанавливаем рабочую директорию для выполнения
WORKDIR /usr/local/bin

# Устанавливаем переменную окружения для конфигурации
ENV CONFIG_WARE_FLOW=/whareflow-server/config/local.yaml

# Открываем порт, если ваше приложение работает на каком-то порте (например, 8080)
EXPOSE 8089

# Запуск приложения
CMD ["wareflow"]
