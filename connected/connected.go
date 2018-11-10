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

package connected

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Connecteddata struct {
	Connected []string `yaml:"connected"`
}

func (conn *Connecteddata) Readconnectvalues() {
	log.Printf("%s", "Connected invoked")

	yamlFile, err := ioutil.ReadFile("/etc/kubedge/connected_ue.yaml")
	//yamlFile, err := ioutil.ReadFile("/tmp/connected_ue.yaml")
	log.Printf("yamlFile=%s", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, conn)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	//(*conn).Connected[0] = "hello world"    //WORKS - change seen in caller
	log.Printf("in library connected: connected=%s", conn.Connected)

}
