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
1. setup env
```
export PORT=<server port>
export MYSQL_DSN=<mysql dsn>
export ALLOW_CORS_ORIGIN=<cors origin>
```

2. command
```sh
$ make run
```

# test
```
$ make test
```