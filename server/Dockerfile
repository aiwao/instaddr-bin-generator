FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.work ./
COPY client ./client
COPY server ./server
COPY common ./common

RUN apk add --no-cache git sqlite-dev gcc musl-dev
RUN cd server && go mod download

WORKDIR /app/server
RUN go build -o server

FROM alpine
RUN apk add --no-cache sqlite-libs
WORKDIR /app/server

COPY --from=builder /app/server/server ./server

CMD ["./server"]