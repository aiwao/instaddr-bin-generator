FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.work ./
COPY client ./client
COPY common ./common
COPY server ./server

RUN apk add --no-cache git sqlite-dev gcc musl-dev
RUN cd client && go mod download

WORKDIR /app/client
RUN go build -o client

FROM alpine
WORKDIR /app/client
COPY --from=builder /app/client/client ./client

CMD ["./client"]