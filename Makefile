build:
	@go build -v -o bin/flamed cmd/flamed/main.go

run:
	@go build -v -o bin/flamed cmd/flamed/main.go
	@./bin/flamed

test:
	@go test ./... -v

cover:
	@go test ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./pkg/pb --go_out=./pkg/pb flamed.proto

push:
	git push
	git push github