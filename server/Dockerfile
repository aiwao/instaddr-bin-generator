FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache git sqlite-dev gcc musl-dev

COPY . .
RUN go build -o main .

FROM alpine
RUN apk update && apk upgrade
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main ./main
COPY addrbin.db ./addrbin.db

CMD ["./main"]
