package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "wqp.com/cricket-center/cmd/client/deliver"
)

func main() {
	conn, err := grpc.Dial("192.168.1.163:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCenterPlatformClient(conn)

	// 发送 EventRequest
	sendEvent(c)

	// 执行 ServerStream RPC
	serverStream(c)
}

func sendEvent(c pb.CenterPlatformClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := c.SendEvent(ctx, &pb.EventRequest{
		ProjectNumber: "123",
		EventCode:     "code",
	})
	if err != nil {
		log.Fatalf("could not send event: %v", err)
	}
	log.Printf("Received EventResponse: %s", res.Message)
}

func serverStream(c pb.CenterPlatformClient) {
	stream, err := c.ServerStream(context.Background(), &pb.ControllRequest{ProjectNumber: "123"})
	if err != nil {
		log.Fatalf("could not start server stream: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error during stream: %v", err)
		}
		log.Printf("Received ControllResponse: %+v", res)
	}
}
