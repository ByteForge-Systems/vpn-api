FROM golang:1.23.6-alpine3.21

# Устанавливаем зависимости для сборки и работы
RUN apk add --no-cache bash

WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Устанавливаем migrate CLI
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /main_api cmd/main.go

# Делаем бинарный файл исполняемым
RUN chmod +x /main_api

# Копируем migrate CLI в /usr/local/bin, чтобы он был доступен
RUN cp /go/bin/migrate /usr/local/bin/migrate

# Открываем порт для приложения
EXPOSE 8080

# По умолчанию запускаем приложение
CMD ["/main_api"]