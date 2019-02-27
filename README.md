# go-pg-mutex

Mutex lock based on PostgreSQL advisory locks for Go

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

[Build-Status-Url]: https://travis-ci.org/vgarvardt/go-oauth2-pg
[Build-Status-Image]: https://travis-ci.org/vgarvardt/go-oauth2-pg.svg?branch=master
[codecov-url]: https://codecov.io/gh/vgarvardt/go-oauth2-pg
[codecov-image]: https://codecov.io/gh/vgarvardt/go-oauth2-pg/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/vgarvardt/go-oauth2-pg
[reportcard-image]: https://goreportcard.com/badge/github.com/vgarvardt/go-oauth2-pg
[godoc-url]: https://godoc.org/github.com/vgarvardt/go-oauth2-pg
[godoc-image]: https://godoc.org/github.com/vgarvardt/go-oauth2-pg?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
