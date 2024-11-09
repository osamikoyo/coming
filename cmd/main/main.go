package main

import (
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	"college-hub-coming/cmd/server"
	"college-hub-coming/proto"
)

func main() {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil{
		loger.Error(err.Error())
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	proto.RegisterCommingMakingServer(s, server.Server{})

	if err := s.Serve(lis); err != nil{
		loger.Error(err.Error())
	}
}
