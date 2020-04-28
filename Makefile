test:
	@go test ./...

cover:
	@go test ./... -coverprofile=cover.out -v

cover-html:
	@go tool cover -html=cover.out

protobuf:
	@protoc -I=./pkg/pb --go_out=./pkg/pb flamed.proto
