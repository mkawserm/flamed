build:
	@CGO_ENABLED=0 go build -v -o bin/flamed cmd/flamed/main.go

test:
	@CGO_ENABLED=0 go test ./...

cover:
	@CGO_ENABLED=0 go test ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./pkg/pb --go_out=./pkg/pb flamed.proto
