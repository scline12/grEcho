package pb

import grpc "google.golang.org/grpc"

//go:generate protoc --go_out "${GOPATH}/src" --go-grpc_out "${GOPATH}/src" --proto_path "${GOPATH}/src/github.com/calculi-corp/proto" pb/buildbreaker.proto pb/codecoverage.proto pb/library.proto pb/license.proto pb/pipelineoverview.proto  pb/sca.proto pb/scaproperties.proto  pb/testmetrics.proto pb/veracode.proto pb/vulnerabilities.proto

// GetDesc returns the service description for registering with consul
func GetDesc() grpc.ServiceDesc {
	return EchoService_ServiceDesc
}
