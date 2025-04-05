FROM golang:1.23.6-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main_api cmd/main.go

EXPOSE 8080
CMD ["/main_api"]
