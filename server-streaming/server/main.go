package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	pb "github.com/shunsukesan/grpc-stream-patterns/server-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type Service struct{}

func (s *Service) Notification(req *pb.NotificationRequest, stream pb.NotificationService_NotificationServer) error {
	log.Println("recieved request")
	var i int32
	var err error
	for i = 0; i < req.Num; i++ {
		message := fmt.Sprintf("notification %d", i)
		if err = stream.Send(&pb.NotificationResponse{
			Message: message,
		}); err != nil {
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
	pb.RegisterNotificationServiceServer(grpcServer, &Service{})
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
