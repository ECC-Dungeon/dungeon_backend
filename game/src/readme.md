# GRPC の設定
Golang のモジュール設定
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

移動してコンパイル
```
sudo apt update
sudo apt install protobuf-compiler
cd ./grpc
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. server.proto
```