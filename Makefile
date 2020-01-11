ARG = ARG
DC = docker-compose
CURRENT_DIR = $(shell pwd)

sql:
	docker run --rm -v $(CURRENT_DIR)/migrate/sql:/sql \
	migrate/migrate:latest create -ext sql -dir /sql ${ARG}

sql-doc:
	docker run --rm --net=api-gateway_default -v $(CURRENT_DIR)/migrate:/work k1low/tbls \
	doc mysql://root:password@chat-db:3306/chat_DB ./

proto:
	docker run --rm -v $(CURRENT_DIR)/src/chat_grpc:$(CURRENT_DIR) -w $(CURRENT_DIR) thethingsindustries/protoc \
	-I. \
	-I/usr/include/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:. \
	--validate_out="lang=go:." \
	--doc_out=markdown,README.md:./ \
	chat.proto

up:
	$(DC) up

build:
	$(DC) build

down:
	$(DC) down