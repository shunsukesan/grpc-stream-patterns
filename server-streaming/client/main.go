package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/shunsukesan/grpc-stream-patterns/server-streaming/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func run(num int32) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	req := &pb.NotificationRequest{
		Num: num,
	}

	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Println("message: ", reply.GetMessage())
	}

	return nil
}

func main() {
	var (
		numFlag = flag.Int("num", 1, "num")
	)
	flag.Parse()

	num := *numFlag

	if err := run(int32(num)); err != nil {
		log.Fatal(err)
	}
}
