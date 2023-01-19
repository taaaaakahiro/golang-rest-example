# golang-rest-example

# my version
```
Docker version 20.10.12, build e91ed57
docker-compose version 1.29.2, build 5becea4c
```

# setup
## isntall golang & docker
   - go
      - https://go.dev/
   - docker desktop
      - https://www.docker.com/products/docker-desktop/

## run DB *DBコンテナを起動
```sh
$ docker-compose up -d #DBコンテナ起動
$ docker-compose down #DBコンテナ停止
```

## run app *api serverを起動
1. setup environment/環境変数を設定(内容は.env.sampleファイル参照)
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
- 画面キャプチャ
   - /version
      - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D--version
   - /healthz
      - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D-healthz

# test *goでテストを実行
```sh
$ make test #go test ./...
```

# curl/http request
```sh
$ curl -X GET localhost:8080/version #バージョン確認(動作確認も兼ねています)
$ curl -X GET localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserを取得
$ curl -X GET localhost:8080/v1/users -H "Content-Type: application/json" #userテーブルの一覧(全件)を取得
$ curl -X POST localhost:8080/v1/user -H "Content-Type: application/json" --data-raw '{"name": "user"}' #usersテーブルに指定したnameのuserレコードを追加
$ curl -X DELETE localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserをテーブルから削除

```
   - 画面キャプチャ
      - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D%E7%94%BB%E9%9D%A2%E3%82%AD%E3%83%A3%E3%83%97%E3%83%81%E3%83%A3

# Architecture/ディレクト構成
## pkg
   - command: 各pkgを初期化
   - config: 設定を管理
   - domain: 目的の明確化
   - handler: ルーティング
   - infrastractue
      - persistence: クエリ
   - io: DB接続
   - middleware: リクエストの前処理等 
   - server: API Serverの設定
   - version: バージョン管理