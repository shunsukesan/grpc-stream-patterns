package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
	pb "github.com/shunsukesan/grpc-stream-patterns/bidirectional-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50003"
)

type server struct{}

func (s *server) Message(stream pb.MessageService_MessageServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		content := req.GetContent()
		log.Println("recieved message: ", content)

		res := fmt.Sprintf("Response: %s", content)
		err = stream.Send(&pb.MessageResponse{Content: res})
		if err != nil {
			return err
		}
	}

	return nil
}

func run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
