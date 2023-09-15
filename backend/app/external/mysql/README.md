# backend

## migrate

ツールのインストール
```sh
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
````

migrate fileの作成
```sh
migrate create -ext sql -dir external/mysql/migrations -seq 名前
```


migrate 実行
```sh
migrate -path パス -database="mysql://ユーザー名:パスワード@ホスト名:ポート番号/データベース名?sslmode=disable" up 1
```

一つ実行
```sh
migrate -path external/mysql/migrations -database="mysql://bottlist:docker@tcp(db:3306)/bottlist?multiStatements=true" up 1
```

全部実行
```sh
migrate -path external/mysql/migrations -database="mysql://bottlist:docker@tcp(db:3306)/bottlist?multiStatements=true" up
```