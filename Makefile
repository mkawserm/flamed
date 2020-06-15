build:
	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/flamed.go

build-darwin-amd64:
	@GOOS=darwin GOARCH=amd64 go build -v -o bin/flamed-darwin-amd64 cmd/flamed/flamed.go

#build-darwin-arm64:
#	@GOOS=darwin GOARCH=arm64 go build -v -o bin/flamed-darwin-arm64 cmd/flamed/flamed.go

build-windows-amd64:
	@GOOS=windows GOARCH=amd64 go build -v -o bin/flamed-windows-amd64.exe cmd/flamed/flamed.go

#build-windows-arm64:
#	@GOOS=windows GOARCH=arm64 go build -v -o bin/flamed-windows-arm64.exe cmd/flamed/flamed.go

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -v -o bin/flamed-linux-amd64 cmd/flamed/flamed.go

build-linux-arm64:
	@GOOS=linux GOARCH=arm64 go build -v -o bin/flamed-linux-arm64 cmd/flamed/flamed.go

build-all: build-darwin-amd64 build-windows-amd64 build-linux-amd64 build-linux-arm64

#	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/flamed.go

run:
	@go run cmd/flamed/flamed.go run --notify-commit true --node-id 1 --storage-path /tmp/data1 --http-server-address 0.0.0.0:8081 --raft-address 0.0.0.0:63001 --log-level debug

run-race:
	@go run -race cmd/flamed/flamed.go run --notify-commit true --node-id 1 --storage-path /tmp/data1 --http-server-address 0.0.0.0:8081 --raft-address 0.0.0.0:63001 --log-level debug

test:
	@go test ./... -v

cover:
	@go test ./... -coverprofile=cover.out -v

#test:
#	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -v
#
#cover:
#	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./ -I=./pkg/pb --go_out=./pkg/pb flamed.proto
	@protoc -I=./ -I=./pkg/tp/intkey --go_out=./pkg/tp/intkey intkey.proto
	@protoc -I=./ -I=./pkg/tp/json --go_out=./pkg/tp/json json.proto

push:
	git push
	git push github