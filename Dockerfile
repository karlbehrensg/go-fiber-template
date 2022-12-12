FROM golang:1.19.2-alpine3.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

# install dependency
RUN go install github.com/swaggo/swag/cmd/swag@latest

# generate swagger docs
RUN swag init -g cmd/api/main.go

# build compilate
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/api/main.go

FROM alpine:3.16.2 AS production

COPY --from=builder /app/bin/main /app/bin/main

ARG	ENV
ARG APP_PORT
ARG	DB_HOST
ARG	DB_PORT
ARG	DB_USER
ARG	DB_PASSWORD
ARG	DB_NAME
ARG	DB_SCHEMA
ARG	DB_SSL_MODE
ARG	DB_TIME_ZONE
ARG	JWT_SECRET
EXPOSE ${APP_PORT}

CMD [ "/app/bin/main" ]
