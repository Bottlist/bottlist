# 変数
# ディレクトリ指定
ifdef DIR
  DIRPATH:=$(shell echo ./`dirname $(DIR)`/`basename $(DIR) /`/)
else
  DIRPATH:=./
endif


# 自動生成コンテナ、MySQLなどのミドルウェアコンテナの立ち上げ
# run_game, run_adminをする場合は必須でdocker_upをする必要あり
.PHONY: docker_up
docker_up:
	docker compose -f docker-compose.gen.yml up --build -d

.PHONY: docker_run
docker_run:
	docker compose -f docker-compose.gen.yml up -d

# docker立ち上げ + protoからGoコードを自動生成
.PHONY: proto
proto:
	docker compose -f docker-compose.gen.yml up -d
	make docker_proto
	docker compose -f docker-compose.gen.yml down

# protoからGoコードを自動生成
.PHONY: docker_proto
docker_proto:
	docker container exec bottlist_gen sh -c '(cd /app && make generate_proto)'

.PHONY: generate_proto
generate_proto:
	sh scripts/gen_proto.sh

# wireによる自動生成
.PHONY: wire
wire:
	docker compose -f docker-compose.gen.yml up -d
	make docker_wire
	docker compose -f docker-compose.gen.yml down

# wireによる自動生成
.PHONY: docker_wire
docker_wire:
	docker container exec bottlist_gen sh -c '(wire cmd/di/wire.go)'

.PHONY: install
install:
	go clean -cache
	go install github.com/gogo/protobuf/protoc-gen-gofast@v1.3.2
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
	go install github.com/google/wire/cmd/wire@v0.5.0
