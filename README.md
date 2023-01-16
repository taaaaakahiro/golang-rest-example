# golang-rest-example

# my version
```
Docker version 20.10.12, build e91ed57
docker-compose version 1.29.2, build 5becea4c
```

# setup


## run DB *DBコンテナを起動
```sh
$ docker-compose up -d #コンテナ起動
$ docker-compose down #コンテナ停止
```

## run app *api serverを起動
1. setup environment/環境変数を設定
```sh
$ export PORT=<server port>
$ export MYSQL_DSN=<mysql dsn>
$ export ALLOW_CORS_ORIGIN=<cors origin>
```
2. command/ローカルでapi serverを起動
```sh
$ make run #go run ./cmd/api/main.go
```
3. check endpoint/ブラウザで動作確認
 - <SERVER PORT>は1.で指定したPORT
```
$ localhost:<SERVER PORT>/version #ex. localhost:8080/version
$ localhost:<SERVER PORT>/healthz
```

# test ~goでテストを実行~
```sh
$ make test #go test ./...
```

# http request
```sh
$ curl -X GET localhost:8080/version #バージョン確認(動作確認も兼ねています)
$ curl -X GET localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserを取得
$ curl -X GET localhost:8080/v1/users -H "Content-Type: application/json" #userテーブルの一覧(全件)を取得
$ curl -X POST localhost:8080/v1/user -H "Content-Type: application/json" --data-raw '{"name": "user"}' #usersテーブルに指定したnameのuserレコードを追加
$ curl -X DELETE localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserをテーブルから削除

```
