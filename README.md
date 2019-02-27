# go-pg-mutex

Mutex lock based on PostgreSQL advisory locks for Go. Can be used to acquire a lock on a resource exclusively across several running instances of the application.

## Install

```bash
$ go get -u -v github.com/vgarvardt/go-pg-mutex
```

## PostgreSQL drivers

The store accepts an adapter interface that interacts with the DB. Adapter and implementations are extracted to separate package [`github.com/vgarvardt/go-pg-adapter`](https://github.com/vgarvardt/go-pg-adapter) for easier maintenance.

## Usage example

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
	"github.com/vgarvardt/go-pg-adapter/pgxadapter"
	"github.com/vgarvardt/go-pg-mutex"
)

func main() {
	pgxConnConfig, _ := pgx.ParseURI(os.Getenv("DB_URI"))
	pgxConn, _ := pgx.Connect(pgxConnConfig)
	
	m, _ := pgmutex.New(pgxadapter.NewConn(pgxConn))
	
	useExclusiveResource(m)
	
	if !useExclusiveResourceOrNoOp(m) {
		fmt.Println("Resource is busy, doing nothing")
	}
}

func useExclusiveResource(m *pgmutex.PgMutex) {
	lockName := "my-lock"
	_ := m.Lock(lockName)
	defer m.Unlock(lockName)

	// do something with resource exclusively across several instances
}

func useExclusiveResourceOrNoOp(m *pgmutex.PgMutex) bool {
	lockName := "my-try-lock"
	if success, _ := m.TryLock(lockName); !success {
		return false
	}
	defer m.Unlock(lockName)

	// do something with resource exclusively across several instances
}
```

## How to run tests

You will need running PostgreSQL instance. E.g. the one running in docker and exposing a port to a host system

```bash
docker run --rm -p 5432:5432 -it -e POSTGRES_PASSWORD=mutex -e POSTGRES_USER=mutex -e POSTGRES_DB=mutex postgres:10
```

Now you can run tests using the running PostgreSQL instance using `PG_URI` environment variable

```bash
PG_URI=postgres://mutex:mutex@localhost:5432/mutex?sslmode=disable go test -cover ./...
```

## MIT License

```
Copyright (c) 2019 Vladimir Garvardt
```

[Build-Status-Url]: https://travis-ci.org/vgarvardt/go-pg-mutex
[Build-Status-Image]: https://travis-ci.org/vgarvardt/go-pg-mutex.svg?branch=master
[codecov-url]: https://codecov.io/gh/vgarvardt/go-pg-mutex
[codecov-image]: https://codecov.io/gh/vgarvardt/go-pg-mutex/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/vgarvardt/go-pg-mutex
[reportcard-image]: https://goreportcard.com/badge/github.com/vgarvardt/go-pg-mutex
[godoc-url]: https://godoc.org/github.com/vgarvardt/go-pg-mutex
[godoc-image]: https://godoc.org/github.com/vgarvardt/go-pg-mutex?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
