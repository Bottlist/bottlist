docker compose は　proto生成

docker内で実行

protoc -I. --go_out=api/go --go-grpc_out=api/go api/proto/api/*.proto