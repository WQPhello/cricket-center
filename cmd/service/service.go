package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "wqp.com/cricket-center/cmd/service/deliver"
)

type server struct {
	pb.UnimplementedCenterPlatformServer
}

func (s *server) SendEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	// 在这里处理 EventRequest
	log.Printf("Received EventRequest: %+v", req)

	// 返回 EventResponse
	return &pb.EventResponse{
		ProjectNumber: req.ProjectNumber,
		EventCode:     req.EventCode,
		Message:       "Processed EventRequest",
	}, nil
}

func (s *server) ServerStream(req *pb.ControllRequest, stream pb.CenterPlatform_ServerStreamServer) error {
	// 在这里处理 ControllRequest
	log.Printf("Received ControllRequest: %+v", req)

	// 模拟发送一系列 ControllResponse
	for i := 0; i < 100; i++ {
		err := stream.Send(&pb.ControllResponse{
			A: &pb.ControllA{A1: "Data", A2: int64(i)},
			B: &pb.ControllB{B1: "More Data", B2: true},
		})
		if err != nil {
			return err
		}

		//
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "192.168.1.163:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCenterPlatformServer(grpcServer, &server{})

	//
	log.Println("Server listening at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
