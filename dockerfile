FROM golang:1.17-alpine AS builder
ADD src/go.* /app/
WORKDIR /app
RUN go mod download


FROM builder as build
RUN go install github.com/cosmtrek/air@v1.27.9
# RUN apk add bash build-base

WORKDIR /app

# CMD ["air", "-c", "cmd/api/.air.toml"]
