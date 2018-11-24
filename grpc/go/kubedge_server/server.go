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

func (s *kubedgeserver) MME_PDN_CONNECT(ctx context.Context, in *pb.MME_PDN_CONNECT_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_PDN_CONNECT with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_PDN_DISCONNECT(ctx context.Context, in *pb.MME_PDN_DISCONNECT_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_PDN_DISCONNECT with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_ATTACH(ctx context.Context, in *pb.MME_ATTACH_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_ATTACH with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_DETACH(ctx context.Context, in *pb.MME_DETACH_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_DETACH with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_HO(ctx context.Context, in *pb.MME_HO_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_HO with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_TAU(ctx context.Context, in *pb.MME_TAU_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_TAU with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) MME_SERVICE_REQUEST(ctx context.Context, in *pb.MME_SERVICE_REQUEST_MSG) (*pb.SIMReply, error) {
	message := "Server received MME_SERVICE_REQUEST with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_SESSION_CREATION(ctx context.Context, in *pb.SGW_SESSION_CREATION_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_SESSION_CREATION with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_SESSION_DELETION(ctx context.Context, in *pb.SGW_SESSION_DELETION_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_SESSION_DELETION with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_BEARER_CREATION(ctx context.Context, in *pb.SGW_BEARER_CREATION_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_BEARER_CREATION with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_BEARER_DELETION(ctx context.Context, in *pb.SGW_BEARER_DELETION_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_BEARER_DELETION with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_BEARER_MODIFICATION(ctx context.Context, in *pb.SGW_BEARER_MODIFICATION_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_BEARER_MODIFICATION with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_DEDICATED_BEARER_ACTIVATE(ctx context.Context, in *pb.SGW_DEDICATED_BEARER_ACTIVATE_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_DEDICATED_BEARER_ACTIVATE with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) SGW_DEDICATED_BEARER_DEACTIVATE(ctx context.Context, in *pb.SGW_DEDICATED_BEARER_DEACTIVATE_MSG) (*pb.SIMReply, error) {
	message := "Server received SGW_DEDICATED_BEARER_DEACTIVATE with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_RRC_CONN_SETUP(ctx context.Context, in *pb.ENB_RRC_CONN_SETUP_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_RRC_CONN_SETUP with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_S1_SIG_CONN_SETUP(ctx context.Context, in *pb.ENB_S1_SIG_CONN_SETUP_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_S1_SIG_CONN_SETUP with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_INITIAL_CTXT_SETUP(ctx context.Context, in *pb.ENB_INITIAL_CTXT_SETUP_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_INITIAL_CTXT_SETUP with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_UE_CTX_RELEASE(ctx context.Context, in *pb.ENB_UE_CTX_RELEASE_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_UE_CTX_RELEASE with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_X2HO_PREP_OUT(ctx context.Context, in *pb.ENB_X2HO_PREP_OUT_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_X2HO_PREP_OUT with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_X2HO_PREP_IN(ctx context.Context, in *pb.ENB_X2HO_PREP_IN_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_X2HO_PREP_IN with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_X2HO_EXEC_OUT(ctx context.Context, in *pb.ENB_X2HO_EXEC_OUT_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_X2HO_EXEC_OUT with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_X2HO_EXEC_IN(ctx context.Context, in *pb.ENB_X2HO_EXEC_IN_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_X2HO_EXEC_IN with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_ERAB_SETUP(ctx context.Context, in *pb.ENB_ERAB_SETUP_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_ERAB_SETUP with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func (s *kubedgeserver) ENB_ERAB_RELEASE(ctx context.Context, in *pb.ENB_ERAB_RELEASE_MSG) (*pb.SIMReply, error) {
	message := "Server received ENB_ERAB_RELEASE with IMSI: " + in.Imsi
	log.Printf(message)
	append2log(message)
	return &pb.SIMReply{Message: message}, nil
}

func Server(simuname string, c config.Configdata) {
	conf = c
	log.Printf("%s: server is running, enable_log=%v", simuname, conf.Enable_log)

	lis, err := net.Listen("tcp", conf.Kubedge_server_port)
	if err != nil {
		log.Fatalf("%s: failed to listen: %v", simuname, err)
	}
	server := grpc.NewServer()
	pb.RegisterKubedgeServer(server, &kubedgeserver{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("%s: failed to serve: %v", simuname, err)
	}
	log.Fatalf("%s: server is exiting", simuname, err)
}
