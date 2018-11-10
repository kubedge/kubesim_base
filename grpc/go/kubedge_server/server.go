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

package server

import (
	"github.com/kubedge/kubesim_base/config"
	pb "github.com/kubedge/kubesim_base/grpc/go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const (
	port        = ":50051"
	logfilename = "/etc/kubedge/location_ue.txt"
)

var conf config.Configdata

type kubedgeserver struct{}

func append2log(message string) {
	if conf.Enable_log {
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
	message := "server received Message: " + in.Protocol + ":: Received data from eNodeB"
	log.Printf(message)
	append2log(message)
	return &pb.EPCReply{Message: in.Protocol + ":: Received data from eNodeB"}, nil
}

func (s *kubedgeserver) DetectNW(ctx context.Context, in *pb.UERequest) (*pb.EPCReply, error) {
	message := "server received Message: " + in.Network + ":: is the mode of connection"
	log.Printf(message)
	append2log(message)
	return &pb.EPCReply{Message: in.Network + ":: is the mode of connection"}, nil
}

func Server(simuname string, c config.Configdata) {
	conf = c
	log.Printf("%s: server is running, enable_log=%s", simuname, conf.Enable_log)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("%s: failed to listen: %v", simuname, err)
	}
	s := grpc.NewServer()
	pb.RegisterKubedgeServer(s, &kubedgeserver{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("%s: failed to serve: %v", simuname, err)
	}
	log.Fatalf("%s: server is exiting", simuname, err)
}
