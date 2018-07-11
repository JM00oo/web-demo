# web-demo
## pre-requirement
### govendor https://github.com/kardianos/govendor
### sql-migrate https://github.com/rubenv/sql-migrate

## install dependency
```
$ govendor sync
```

## modify dbconfig.yml and update schema, ref: dbconfig-example.yml
```
$ vim ./dbconfig.yml
$ sql-migrate up
```
## setup SQL Config, ref: config/config-exapmle.go
```
$ vim ./config/config.go
```
## Run test
```
$ GO_ENV=unit-test go test ./...
```

## Run server
```
$ go run main.go
```
> Visit localhost:8080
