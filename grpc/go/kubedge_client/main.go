package main

import (
        "log"
        "os"
        "time"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
	pb  "github.com/kubedge/kubesim_base/grpc/go"
)

const (
	address     = "172.17.0.2:50051"
        //address     = "localhost:50051"
        defaultNetwork = "Bluetooth"
        S1ProtoHead = "X1000001"
        SNProtoHead = "XN888888"
	defaultProtocol =  "S1"
	defaultDemo = "protocol"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
            log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewKubedgeClient(conn)

    // Contact the server and print out its response.
    demo := defaultDemo
    network := defaultNetwork
    protocol := defaultProtocol
    protoHead := S1ProtoHead
    if len(os.Args) > 1 {
       demo = os.Args[1]
       if (demo == defaultDemo) {
           protocol = os.Args[2]
           if protocol != defaultProtocol {
               protoHead =  SNProtoHead
           }
       } else {
           network = os.Args[2]
       }
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    if demo == defaultDemo {
        r, err := c.FiveGDemo(ctx, &pb.EnodeRequest{Protocol: protoHead})
        if err != nil {
            log.Fatalf("Error: %v", err)
        }
        log.Printf("%s", r.Message)
    } else {
        rr, errr := c.DetectNW(ctx, &pb.UERequest{Network: network})
        if errr != nil {
            log.Fatalf("Error: %v", errr)
        }
        log.Printf("%s", rr.Message)
    }
}


