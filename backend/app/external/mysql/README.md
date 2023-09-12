# backend

## migrate

```sh
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
````

docker内で行うコマンド
```sh
migrate -path パス -database="mysql://ユーザー名:パスワード@ホスト名:ポート番号/データベース名?sslmode=disable" up 1
```

```sh
migrate -path external/mysql/migrations -database="mysql://bottlist:docker@tcp(db:3306)/bottlist?multiStatements=true" up 1
```
