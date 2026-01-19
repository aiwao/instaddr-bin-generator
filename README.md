# instaddr-bin-generator
## Instaddr Account generator/database with Docker + Go + SQLite3

## ENV
```
CREATE_ACCOUNT_DELAY: Delay for create a account (default: 1000) (ms)
CREATE_ADDRESS_DELAY: Delay for create a address (default: 1000) (ms)
ON_ERROR_DELAY: Delay for sleep on error (default: 5000) (ms)
ADDRESS_AMOUNT: Maximum amount of create address in account (50)
MUST_LEGIT_TO_AMOUNT: No skipping address creation by error (default: true)
PROXY: Network proxy for instaddr api (recommend use oxylabs)
```

## RUN
```shell
docker compose --env-file .env.local up app
```

## Connect to the database (example)
```shell
psql --host localhost --port 5432 --username user --password password --dbname accountdb
```

## Get accounts data as json array
```shell
psql --host localhost --port 5432 --username user --password password --dbname accountdb -t -A -c "SELECT json_agg(row_to_json(t)) FROM (SELECT * FROM accounts) t;" > accounts.json
```
