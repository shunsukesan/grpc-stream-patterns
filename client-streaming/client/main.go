package main

import (
	"context"
	"io"
	"os"

	"github.com/labstack/gommon/log"
	pb "github.com/shunsukesan/grpc-stream-patterns/client-streaming/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50002"
)

func run() error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewVideoServiceClient(conn)

	stream, err := client.Upload(context.Background())
	if err != nil {
		return err
	}

	// hardcoded
	file, err := os.Open("./sample.mp4")
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		stream.Send(&pb.UploadRequest{VideoData: buf})
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
