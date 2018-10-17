package server

import (
        "log"
        "net"
        "os"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
	"github.com/kubedge/kubesim_base/config"
	pb "github.com/kubedge/kubesim_base/grpc/go"
)

const (
        port = ":50051"
        logfilename = "/etc/kubedge/location_ue.txt"
)

var conf config.Configdata

type kubedgeserver struct{}

func append2log(message string) {
        if (conf.Enable_log) {
            f, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
            if err == nil {
                f.WriteString(message + "\n") 
                f.Close()
            } else {
	        log.Printf("error opening logfile %v", err)
            }
        }
}

func (s *kubedgeserver) FiveGDemo(ctx context.Context, in *pb.EnodeRequest) (*pb.EPCReply, error) {
	message := "server recieved Message: " + in.Protocol +  ":: Recieved data from Enode"
	log.Printf(message)
	append2log(message)
	return &pb.EPCReply{Message: in.Protocol +  ":: Recieved data from Enode"}, nil
}

func (s *kubedgeserver) DetectNW(ctx context.Context, in *pb.UERequest) (*pb.EPCReply, error) {
	message := "server received Message: " + in.Network +  ":: is the mode of connection"
	log.Printf(message)
	append2log(message)
	return &pb.EPCReply{Message: in.Network +  ":: is the mode of connection"}, nil
}

func Server(c config.Configdata) {
        conf = c
        log.Printf("server is running, enable_log=%s", conf.Enable_log)
        //append2log("test write to logfile") 

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
        log.Fatalf("server is exiting", err)
}

