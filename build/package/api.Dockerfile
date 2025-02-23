FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/api/ ./cmd/api/
COPY internal/ ./internal/
COPY pkg/ ./pkg/

RUN go build -o /api ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api .
COPY .env .env

ENV $(cat .env | xargs)

CMD ["/root/api"]
