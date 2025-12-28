# instaddr-bin-generator
## Instaddr Account generator/database with Docker + Go + SQLite3

## Run

### Server
#### ENV
```
CREATE_ACCOUNT_DELAY: Delay for create a account (default: 1000) (ms)
CREATE_ADDRESS_DELAY: Delay for create a address (default: 1000) (ms)
ON_ERROR_DELAY: Delay for sleep on error (default: 5000) (ms)
ADDRESS_AMOUNT: Maximum amount of create address in account (50)
MUST_LEGIT_TO_AMOUNT: No skipping address creation by error (default: 0) (0: false, 1: true)
PROXY: Network proxy for instaddr api (recommend use oxylabs)
```

```
docker compose up --build bin-gen-server
```

### Client
#### ENV
```
Local: Use local database (default: 0) (true/false)
SERVER_URL: Custom server address (default: http://localhost:8080)
AMOUNT_ACCOUNT: Account amount to get (default: 100)
MIN_AMOUNT_ADDRESS: Minimum amount of addresses in account (default: 10)
```

```
docker compose up --build bin-gen-client
```