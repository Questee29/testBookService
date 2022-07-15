package grpcServer

import (
	"net"

	pb "github.com/Questee29/testBookService/proto/protob"
	"google.golang.org/grpc"
)

type Deps struct {
	GrpcHandler pb.BooksServiceServer
}
type Server struct {
	deps Deps
	srv  *grpc.Server
}

func NewServer(deps Deps) *Server {
	return &Server{
		srv:  grpc.NewServer(),
		deps: deps,
	}
}

func (s *Server) ListenAndServe(port string) error {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	pb.RegisterBooksServiceServer(s.srv, s.deps.GrpcHandler)
	if err := s.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
