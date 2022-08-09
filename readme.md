![logo](docs/logo.svg)

[![](https://img.shields.io/badge/-Swagger%20Docs-informational?style=flat&logo=swagger&color=blue&labelColor=gray)](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

# REST API for ready User balance integration

- [x] One-tap deployment with [docker-compose](#deployment)
- [x] Exchange rates via [redis]()
- [x] Fast Caching 
- [x] Swagger API documentation + Client generation
- [x] Pprof, Prometheus metrics
- [x] PostgreSQL
- [x] Tests integration with [Allure](https://www.allure.com/)
- [ ] gRPC
- [ ] Rewritten in Rust

---

## API features:

- [x] Get/Change balance
- [x] Transfer between users
- [x] View transaction history
- [x] Block/Unblock balance
- [ ] Taxes
- [ ] Actions with blocked user 

---

## Deployment

### docker-compose

Service starts **immediately** after containers are up

```shell
make compose # docker-compose up
```

#### Redis `:6380` **Database** `0`

#### PostgreSQL `:5430`  **User/Password** `postgres` **Database** `balance_api`

---

**Down** containers (volumes won't be destroyed)

```shell
make compose-down # docker-compose down
```

**Debug** build and run

1) Generate swagger docs with [`swag`](https://github.com/swaggo/swag)
2) Rebuild service image

```shell
make compose-debug 
```

---

## Application endpoints

### Default port: `8080`

- #### `/api/` - REST API
- #### `/swagger` - [Swagger API](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml) documentation
- #### `/metrics` - [Prometheus](https://github.com/prometheus/client_golang) metrics
- #### `/debug/pprof/` - [pprof](https://pkg.go.dev/runtime/pprof)

--- 

## Docs

[![](https://img.shields.io/badge/-Swagger%20Docs-informational?style=for-the-badge&logo=swagger&color=blue&labelColor=gray)](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

#### - Run application and open [`/swagger`](http://0.0.0.0:8080/swagger) page

#### - Visit online [API documentation](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

---

### gRPC: soon or never

---

## Logs

```shell
# Terminal (stdout)
15/07 09:58:22 | INFO | app/handler.go:19 	| Initializing storages
15/07 09:58:22 | INFO | app/handler.go:37 	| Initializing handlers
15/07 09:58:22 | INFO | app/handler.go:63 	| Swagger enabled {"endpoint": "/swagger"}
15/07 09:58:22 | INFO | app/handler.go:75 	| Pprof enabled {"endpoint": "/debug/pprof"}
15/07 09:58:22 | INFO | app/handler.go:82 	| Prometheus metrics enabled {"endpoint": "/metrics"}
15/07 09:58:22 | INFO | app/listen.go:45 	| Server started {"address": "0.0.0.0:8080", "https": false}```
```

```shell
# File (default log.txt)
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/handler.go:19","msg":"Initializing storages"}
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/handler.go:37","msg":"Initializing handlers"}
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/handler.go:63","msg":"Swagger enabled","endpoint":"/swagger"}
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/handler.go:75","msg":"Pprof enabled","endpoint":"/debug/pprof"}
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/handler.go:82","msg":"Prometheus metrics enabled","endpoint":"/metrics"}
{"level":"info","ts":"Fri, 15 Jul 2022 09:58:22 UTC","caller":"app/listen.go:45","msg":"Server started","address":"0.0.0.0:8080","https":false}```
```

--- 

## Flags

#### docker-compose

```yaml
app:
  command:
    - "-pprof" # pprof
    - "-swagger" # swagger docs
    - "-prom" # prometheus docs
    - "-https"
```

#### Standalone run

With non-standard log file path:

```shell
app -log=log.txt
```

HTTPS

```shell
app -https
```

Docs and Metrics

```shell
app -swagger -prom -pprof
```

---

## Config

### All options are loaded from **[.env](.env)**

```dotenv
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
...
REDIS_HASHMAP=currency
...
HOST_ADDRESS=0.0.0.0
HOST_PORT=8080
```

#### docker-compose overrides database credentials

---

## Tests

#### SOME TEST WON'T PASS WITHOUT WORKING APPLICATION

```
go test ./...
```

#### Allure data: [`/tests/allure-results`](`/tests/allure-results`)

---

## Building and Running

#### 1. Clone repository

    `git clone https://github.com/illiafox/balance-api`

#### 2. Setup [env file](.env)

#### 3. Build and Run

```shell
make build
make run # ./app
```

## PostgreSQL Migrations

```shell
migrate -database ${POSTGRESQL_URL} -path migrate/ up
```

## Redis

#### Example: Store currencies in Hash

```shell
HSET currency EUR 65 
```

Where `currency` is `Redis Hash` Name

---


