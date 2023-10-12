gen-call:
	protoc --go_out=. --go_opt=paths=source_relative proto/chatapp.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/chatapp.proto
run-server:
	go run server/main.go
run-client:
	go run client/main.go