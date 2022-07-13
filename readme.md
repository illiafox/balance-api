![logo](docs/logo.svg)

[![](https://img.shields.io/badge/-Swagger%20Docs-informational?style=flat&logo=swagger&color=blue&labelColor=gray)](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

# REST API for ready User balance integration

- [x] One-tap deployment with [docker-compose]()
- [x] Exchange rates via [redis]()
- [x] Swagger API documentation + Client generation
- [x] Pprof, Prometheus metrics
- [x] PostgreSQL
- [ ] gRPC
- [ ] Rewrite in Rust

---

## docker-compose

Service starts **immediately** after containers are up

```shell
make compose # docker-compose up
```

**Down** containers (volumes won't be destroyed)

```shell
make compose-down # docker-compose down
```

**Debug** build and run
1) Generate swagger docs with [`swag`](https://github.com/swaggo/swag):
```make swagger```
2) Rebuild service image
```shell
make compose-debug 
```

## Application endpoints
### Default port: `8080`
- #### `/api/` - REST API
- #### [`/swagger`](http://0.0.0.0:8080/swagger/) - [Swagger API](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml) documentation
- #### [`/metrics`](http://0.0.0.0:8080/metrics/) - Prometheus metrics
- #### [`/debug/pprof`](http://0.0.0.0:8080/metrics/) - pprof

### Docker-compose
- #### Redis `:6380` **Database** `0`
- #### PostgreSQL `:5430`  **User/Password** `postgres` **Database** `balance_api`

 
--- 

## Docs

### 1. Swagger

[![](https://img.shields.io/badge/-Swagger%20Docs-informational?style=for-the-badge&logo=swagger&color=blue&labelColor=gray)](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

#### - Run application and open [`0.0.0.0:8080/swagger`](http://0.0.0.0:8080/swagger) page

#### - Visit online [API documentation](https://validator.swagger.io/?url=https://raw.githubusercontent.com/illiafox/balance-api/master/docs/swagger.yaml)

### 2. gRPC

#### Soon...

---

## Logs

```shell
# Terminal (stdout)
27/06 10:49:30 | INFO | app/handler.go:15 | Initializing storages
27/06 10:49:30 | INFO | app/handler.go:34 | Initializing handlers
27/06 10:49:30 | INFO | app/listen.go:44 | Server started {"address": "0.0.0.0:8080", "https": false}
```

```shell
# File (default log.txt)
{"level":"info","ts":"Mon, 27 Jun 2022 10:49:30 UTC","caller":"app/handler.go:15","msg":"Initializing storages"}
{"level":"info","ts":"Mon, 27 Jun 2022 10:49:30 UTC","caller":"app/handler.go:34","msg":"Initializing handlers"}
{"level":"info","ts":"Mon, 27 Jun 2022 10:49:30 UTC","caller":"app/listen.go:44","msg":"Server started","address":"0.0.0.0:8080","https":false}
```

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

### With non-standard log file path:

```shell
app -log=log.txt
```

### HTTPS

```shell
app -https
```

### Disable swagger

```shell
app -swagger
```

## PostgreSQL Migrations

```shell
migrate -database ${POSTGRESQL_URL} -path migrate/ up
```

## Redis

### Example: Store currencies in Hash

```shell
HSET currency EUR 65 
```

Where `currency` is `Redis Hash` Name (`Redis.HashMap`)

---


