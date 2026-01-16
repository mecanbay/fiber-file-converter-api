

FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/main.go


FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates



COPY --from=builder /app/app /app/cmd/app


COPY config /app/config


RUN mkdir -p /app/logs
RUN ls -la
EXPOSE 6565
WORKDIR /app/cmd
ENTRYPOINT ["./app"]
