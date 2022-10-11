package echo

//go:generate protoc --go_out "${GOPATH}/src" --go-grpc_out "${GOPATH}/src" --proto_path "${GOPATH}/src/github.com/scline12/grEcho/proto" echo.proto
