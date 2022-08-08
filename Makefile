.PHONE: protoc_plugins, longrunning

protoc_plugins:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

OUT_PATH=./api
longrunning:
	# protoc -I . -I ./third_party \
	# 	--go_out=paths=source_relative:$(OUT_PATH) \
	# 	google/longrunning/operations.proto
	# protoc -I . -I ./third_party \
	# 	--go-grpc_out $(OUT_PATH) --go-grpc_opt paths=source_relative \
	# 	google/longrunning/operations.proto
	protoc -I . -I ./third_party \
		--grpc-gateway_out $(OUT_PATH) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt standalone=true \
		google/longrunning/operations.proto