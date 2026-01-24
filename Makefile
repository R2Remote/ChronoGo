.PHONY: proto
proto:
	export PATH="$$PATH:$(shell go env GOPATH)/bin:/usr/local/bin:/opt/homebrew/bin" && \
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	api/proto/*.proto

# 如果你想一键清理生成文件
.PHONY: clean
clean:
	rm -f api/proto/*.pb.go