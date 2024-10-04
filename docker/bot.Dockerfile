FROM golang:1.23-alpine AS build

WORKDIR /src

COPY . .

RUN go build -o ./bot ./cmd/bot/

FROM alpine:3.20.3

WORKDIR /app

COPY --from=build /src/bot ./bot

CMD ["/app/bot"]