/*
Copyright 2018 Kubedge

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	pb "github.com/kubedge/kubesim_base/grpc/go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "192.168.2.101:30180"
	//address     = "localhost:50051"
	defaultNetwork  = "Bluetooth"
	S1ProtoHead     = "X1000001"
	SNProtoHead     = "XN888888"
	defaultProtocol = "S1"
	defaultDemo     = "protocol"
)

func Client(simuname string, demotype string, demovalue string) {
	log.Printf("%s: %s", simuname, "GRPC Client is starting")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s: did not connect: %v", simuname, err)
	}
	defer conn.Close()
	c := pb.NewKubedgeClient(conn)

	// Contact the server and print out its response.
	demo := defaultDemo
	network := defaultNetwork
	protocol := defaultProtocol
	protoHead := S1ProtoHead
	demo = demotype
	if demo == defaultDemo {
		protocol = demovalue
		if protocol != defaultProtocol {
			protoHead = SNProtoHead
		}
	} else {
		network = demovalue
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if demo == defaultDemo {
		r, err := c.FiveGDemo(ctx, &pb.EnodeRequest{Protocol: protoHead})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: %s", simuname, r.Message)
	} else {
		rr, errr := c.DetectNW(ctx, &pb.UERequest{Network: network})
		if errr != nil {
			log.Fatalf("%s: Error: %v", simuname, errr)
		}
		log.Printf("%s: %s", simuname, rr.Message)
	}
	log.Printf("%s: %s", simuname, "GRPC Client is exiting")
}
