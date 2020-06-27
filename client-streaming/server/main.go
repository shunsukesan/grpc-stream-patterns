package main

import (
	"io"
	"log"
	"net"
	"path/filepath"

	"github.com/golang/go/src/os"
	"github.com/pkg/errors"
	pb "github.com/shunsukesan/grpc-stream-patterns/client-streaming/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50002"
)

type server struct{}

func (s *server) Upload(stream pb.VideoService_UploadServer) error {
	err := os.MkdirAll("tmp", 0777)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join("tmp", "a.mp4"))
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		file.Write(req.VideoData)
	}

	res := &pb.UploadResponse{}
	err = stream.SendAndClose(res)
	if err != nil {
		return err
	}

	return nil
}

func run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVideoServiceServer(grpcServer, &server{})
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
