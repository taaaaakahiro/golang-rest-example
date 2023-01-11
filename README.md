# golang-rest-example

# my version
```
Docker version 20.10.12, build e91ed57
docker-compose version 1.29.2, build 5becea4c
```

## setup


# run DB
```sh
$ docker-compose up -d # run
$ docker-compose down # down
```

# run app
1. setup environment
```
export PORT=<server port>
export MYSQL_DSN=<mysql dsn>
export ALLOW_CORS_ORIGIN=<cors origin>
```
2. command
```sh
$ make run
```
3. check endpoint
```
localhost:<SERVER PORT>/version
localhost:<SERVER PORT>/healthz
```

# test
```
$ make test
```

# http request
```
$ curl -X GET localhost:8080/version
$ curl -X GET localhost:8080/v1/user/{id} -H "Content-Type: application/json"
$ curl -X GET localhost:8080/v1/users -H "Content-Type: application/json"
$ curl -X POST localhost:8080/v1/user -H "Content-Type: application/json" --data-raw '{"name": "user"}'

```
