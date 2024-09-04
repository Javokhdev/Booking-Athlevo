CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://sayyidmuhammad:root@localhost:5432/postgres?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://sayyidmuhammad:root@localhost:5432/postgres?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq booking

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

swag-init:
	swag init -g api/router.go -o api/docs

prot-exp:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

gen-proto:
	protoc --go_out=./ \
    --go-grpc_out=./ \
	protos/*.proto