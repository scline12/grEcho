package echo

import (
	"context"
	"fmt"

	client "github.com/calculi-corp/grpc-client"
	handler "github.com/calculi-corp/grpc-handler"
	"github.com/calculi-corp/log"
	"github.com/scline12/grEcho/pb"
)

const (
	ServiceName = "echo"
)

type EchoHandler struct {
	pb.UnimplementedEchoServiceServer
	metrics *handler.Map
	clt     client.GrpcClient
}

func NewEchoHandler(clt client.GrpcClient) *EchoHandler {
	log.Debug("echo.NewEchoHandler")

	return &EchoHandler{
		metrics: handler.NewMap("echo"),
		clt:     clt,
	}
}

func (eh *EchoHandler) Description() *handler.ServiceDesc {
	return &handler.ServiceDesc{Name: ServiceName, ProtoDesc: pb.GetDesc()}
}

func (eh *EchoHandler) MetricMap() *handler.Map {
	return eh.metrics
}

func (eh *EchoHandler) Healthy() error {
	return nil
}

func (eh *EchoHandler) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Debugf("echo.Echo: incoming msg: %s", req.Message)
	if req.Message == "error" {
		return &pb.EchoResponse{Message: "ERROR"}, fmt.Errorf("requested error")
	}
	return &pb.EchoResponse{Message: req.Message}, nil
}
