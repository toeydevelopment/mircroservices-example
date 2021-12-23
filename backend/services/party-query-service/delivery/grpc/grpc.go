package grpc

import (
	"net"

	"github.com/toeydevelopment/microservices-example/party-query-service/pb/party"
	"github.com/toeydevelopment/microservices-example/party-query-service/usecase"
	"google.golang.org/grpc"
)

func NewGRPCServer(addr string, usc usecase.IUsecase) error {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	h := newHandler(usc)

	party.RegisterPartyQueryServiceServer(server, &h)

	return server.Serve(lis)
}
