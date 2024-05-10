package main

import (
	"context"
	pb "grpctest/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var userid_count = 0
var users = []*pb.User{}

type Server struct {
	pb.UnimplementedUserManagerServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	userid_count = userid_count + 1
	users = append(users, &pb.User{
		Id:   int64(userid_count),
		Name: in.GetName(),
		Age:  in.GetAge(),
		Type: in.GetType(),
	})

	return &pb.CreateUserReply{
		Id: int64(userid_count),
	}, nil
}

func (s *Server) SearchUser(ctx context.Context, in *pb.SearchUserRequest) (*pb.SearchUserReply, error) {
	return &pb.SearchUserReply{
		Users: users,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagerServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
