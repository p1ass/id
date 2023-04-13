.PHONY:lint
lint:
	buf lint


.PHONY:format
format:
	buf format -w


.PHONY:generate
generate:
	buf generate


.PHONY:deps
deps:
	brew install bufbuild/buf/buf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	npm install --global @bufbuild/protoc-gen-connect-es @bufbuild/protoc-gen-es
