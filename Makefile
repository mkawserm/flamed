build:
	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/main.go

run:
	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/main.go
	@./bin/flamed

test:
	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -v

cover:
	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./pkg/pb --go_out=./pkg/pb flamed.proto
	@protoc -I=./pkg/tp/identity --go_out=./pkg/tp/identity identity.proto
	@protoc -I=./pkg/tp/index --go_out=./pkg/tp/index index.proto
	@protoc -I=./pkg/tp/intkey --go_out=./pkg/tp/intkey intkey.proto

push:
	git push
	git push github