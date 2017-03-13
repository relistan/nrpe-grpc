package main

import (
	"log"
	"net"

	"github.com/envimate/nrpe"
	"github.com/relistan/nrpe-grpc/nrperpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = ":50000"
)

type server struct{}

func (s *server) NrpeCheck(ctx context.Context, request *nrperpc.NrpeRequest) (*nrperpc.NrpeReply, error) {
	// TODO ignoring context

	conn, err := net.Dial("tcp", "docker1:5666")
	if err != nil {
		return nil, err
	}

	command := nrpe.NewCommand(request.Name)

	// ssl = true, timeout = 0
	result, err := nrpe.Run(conn, command, true, 0)
	if err != nil {
		return nil, err
	}

	reply := &nrperpc.NrpeReply{
		StatusCode: int32(result.StatusCode),
		StatusLine: result.StatusLine,
	}
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	nrperpc.RegisterCheckServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
