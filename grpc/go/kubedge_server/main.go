package main

import (
        "log"
        "net"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
	pb  "github.com/kubedge/kubesim_base/grpc/go"
)

const (
        port = ":50051"
)


type kubedgeserver struct{}

func (s *kubedgeserver) FiveGDemo(ctx context.Context, in *pb.EnodeRequest) (*pb.EPCReply, error) {
	return &pb.EPCReply{Message: in.Protocol +  ":: Recieved data from Enode"}, nil
}

func (s *kubedgeserver) DetectNW(ctx context.Context, in *pb.UERequest) (*pb.EPCReply, error) {
	return &pb.EPCReply{Message: in.Network +  ":: is the mode of connection"}, nil
}

func main() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterKubedgeServer(s, &kubedgeserver{})
        reflection.Register(s)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}

