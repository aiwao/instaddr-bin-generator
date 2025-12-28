FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.work ./
COPY server ./server
COPY common ./common
COPY client ./client

RUN cd server && go mod download
RUN apk add --no-cache git sqlite-dev gcc musl-dev

COPY . .
WORKDIR /app/server
RUN go build -o server

FROM alpine
RUN apk update && apk upgrade
RUN mkdir /app
WORKDIR /app/server

COPY --from=builder /app/server ./server
COPY /server/addrbin.db ./server/addrbin.db

CMD ["./server"]
