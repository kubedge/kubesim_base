package main

import (
	"log"

	loads "github.com/go-openapi/loads"
	"github.com/kubedge/kubesim_base/health/restapi"
	"github.com/kubedge/kubesim_base/health/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCfgapisrvAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureFlags()
	server.Port = 8080

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
