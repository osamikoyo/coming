package server

import (
	"context"

	pb "college-hub-coming/proto"
)

type Server struct {
	pb.UnimplementedCommingMakingServer
}

func (s Server) Create(ctx context.Context, req *pb.ReqComming) (*pb.Response, error){

}

func (s Server) Get(ctx context.Context, user *pb.UserReq) (*pb.RespComming, error){

}

func (s Server) Remove(ctx context.Context, req *pb.ReqComming) (*pb.Response, error){

}