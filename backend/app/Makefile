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

# protoからGoコードを自動生成
.PHONY: docker_generate_proto
docker_generate_proto:
	docker container exec bottlist_gen sh -c '(cd /app && make generate_proto)'


.PHONY: generate_proto
generate_proto:
	sh scripts/gen_proto.sh

.PHONY: install
install:
	go clean -cache
	go install github.com/gogo/protobuf/protoc-gen-gofast@v1.3.2
	go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
