build:
	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/flamed.go

run:
	@DRAGONBOAT_LOGDB=pebble go build -v -tags dragonboat_no_rocksdb -o bin/flamed cmd/flamed/flamed.go
	@./bin/flamed author

test:
	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -v

cover:
	@DRAGONBOAT_LOGDB=pebble go test -tags dragonboat_no_rocksdb ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./ -I=./pkg/pb --go_out=./pkg/pb flamed.proto
	@protoc -I=./ -I=./pkg/tp/intkey --go_out=./pkg/tp/intkey intkey.proto
	@protoc -I=./ -I=./pkg/tp/json --go_out=./pkg/tp/json json.proto

push:
	git push
	git push github