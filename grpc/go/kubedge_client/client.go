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
	// we have a problem here for 5G Core
	address = "sim-epc-kubesim-epc.default.svc.cluster.local:30180"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if demotype == defaultDemo {
	        protocol := defaultProtocol
	        protoHead := S1ProtoHead
		protocol = demovalue
		if protocol != defaultProtocol {
			protoHead = SNProtoHead
		}
		r, err := c.FiveGDemo(ctx, &pb.EnodeRequest{Protocol: protoHead})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: %s", simuname, r.Message)
	} else if demotype == "MME_PDN_CONNECT" {
		r, err := c.MME_PDN_CONNECT(ctx, &pb.MME_PDN_CONNECT_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_PDN_DISCONNECT" {
		r, err := c.MME_PDN_DISCONNECT(ctx, &pb.MME_PDN_DISCONNECT_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_ATTACH" {
		r, err := c.MME_ATTACH(ctx, &pb.MME_ATTACH_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_DETACH" {
		r, err := c.MME_DETACH(ctx, &pb.MME_DETACH_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_HO" {
		r, err := c.MME_HO(ctx, &pb.MME_HO_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_TAU" {
		r, err := c.MME_TAU(ctx, &pb.MME_TAU_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "MME_SERVICE_REQUEST" {
		r, err := c.MME_SERVICE_REQUEST(ctx, &pb.MME_SERVICE_REQUEST_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_SESSION_CREATION" {
		r, err := c.SGW_SESSION_CREATION(ctx, &pb.SGW_SESSION_CREATION_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_SESSION_DELETION" {
		r, err := c.SGW_SESSION_DELETION(ctx, &pb.SGW_SESSION_DELETION_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_BEARER_CREATION" {
		r, err := c.SGW_BEARER_CREATION(ctx, &pb.SGW_BEARER_CREATION_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_BEARER_DELETION" {
		r, err := c.SGW_BEARER_DELETION(ctx, &pb.SGW_BEARER_DELETION_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_BEARER_MODIFICATION" {
		r, err := c.SGW_BEARER_MODIFICATION(ctx, &pb.SGW_BEARER_MODIFICATION_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_DEDICATED_BEARER_ACTIVATE" {
		r, err := c.SGW_DEDICATED_BEARER_ACTIVATE(ctx, &pb.SGW_DEDICATED_BEARER_ACTIVATE_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "SGW_DEDICATED_BEARER_DEACTIVATE" {
		r, err := c.SGW_DEDICATED_BEARER_DEACTIVATE(ctx, &pb.SGW_DEDICATED_BEARER_DEACTIVATE_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_RRC_CONN_SETUP" {
		r, err := c.ENB_RRC_CONN_SETUP(ctx, &pb.ENB_RRC_CONN_SETUP_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_S1_SIG_CONN_SETUP" {
		r, err := c.ENB_S1_SIG_CONN_SETUP(ctx, &pb.ENB_S1_SIG_CONN_SETUP_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_INITIAL_CTXT_SETUP" {
		r, err := c.ENB_INITIAL_CTXT_SETUP(ctx, &pb.ENB_INITIAL_CTXT_SETUP_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_UE_CTX_RELEASE" {
		r, err := c.ENB_UE_CTX_RELEASE(ctx, &pb.ENB_UE_CTX_RELEASE_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_X2HO_PREP_OUT" {
		r, err := c.ENB_X2HO_PREP_OUT(ctx, &pb.ENB_X2HO_PREP_OUT_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_X2HO_PREP_IN" {
		r, err := c.ENB_X2HO_PREP_IN(ctx, &pb.ENB_X2HO_PREP_IN_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_X2HO_EXEC_OUT" {
		r, err := c.ENB_X2HO_EXEC_OUT(ctx, &pb.ENB_X2HO_EXEC_OUT_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_X2HO_EXEC_IN" {
		r, err := c.ENB_X2HO_EXEC_IN(ctx, &pb.ENB_X2HO_EXEC_IN_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_ERAB_SETUP" {
		r, err := c.ENB_ERAB_SETUP(ctx, &pb.ENB_ERAB_SETUP_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else if demotype == "ENB_ERAB_RELEASE" {
		r, err := c.ENB_ERAB_RELEASE(ctx, &pb.ENB_ERAB_RELEASE_MSG{Imsi: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: Response: %s", simuname, r.Message)
	} else {
		r, err := c.DetectNW(ctx, &pb.UERequest{Network: demovalue})
		if err != nil {
			log.Fatalf("%s: Error: %v", simuname, err)
		}
		log.Printf("%s: %s", simuname, r.Message)
	}
	log.Printf("%s: %s", simuname, "GRPC Client is exiting")
}
