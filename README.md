# instaddr-bin-generator
## Instaddr Account generator/database with Docker + Go + SQLite3

## Run

### Server
```docker compose up --build bin-gen-server```

### Client
#### Args
```
-d: local database file path
-s: custom server address (default: localhost:8080)
-acc: maximum account amount to get (default: 1)
-addr: minimum amount of addresses in account (default: 0)
```

```
cd client
go build
./client -acc 100 -addr 10
```