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
        address     = "localhost:50051"
        defaultNetwork = "Bluetooth"
        defaultProtocol = "X1000001"
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
        network := defaultNetwork
        if len(os.Args) > 1 {
                network = os.Args[1]
        }
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.FiveGDemo(ctx, &pb.EnodeRequest{Protocol: defaultProtocol})
        if err != nil {
                log.Fatalf("Error: %v", err)
        }
        log.Printf("%s", r.Message)
        rr, errr := c.DetectNW(ctx, &pb.UERequest{Network: network})
        if errr != nil {
                log.Fatalf("Error: %v", errr)
        }
        log.Printf("%s", rr.Message)
}


