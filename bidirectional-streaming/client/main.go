package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/shunsukesan/grpc-stream-patterns/bidirectional-streaming/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50003"
)

func request(stream pb.MessageService_MessageClient) error {
	err := stream.Send(&pb.MessageRequest{Content: fmt.Sprintf("message")})
	if err != nil {
		return err
	}

	return nil
}

func recieve(stream pb.MessageService_MessageClient) error {
	done := make(chan struct{})
	errChan := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(done)
				break
			}
			if err != nil {
				errChan <- err
				close(errChan)
			}

			log.Printf("response：%s", res.Content)

			time.Sleep(time.Duration(1000000000))

			err = stream.Send(&pb.MessageRequest{
				Content: time.Now().Format("2006-01-02 15:04:05"),
			})
			if err != nil {
				errChan <- err
				close(errChan)
			}
		}
	}()

	select {
	case <-done:
		return nil
	case err := <-errChan:
		return err
	}
}

func run() error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	stream, err := client.Message(context.Background())
	if err != nil {
		return err
	}

	err = request(stream)
	if err != nil {
		return err
	}

	err = recieve(stream)
	if err != nil {
		return err
	}

	err = stream.CloseSend()
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
