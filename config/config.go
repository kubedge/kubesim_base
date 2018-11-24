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

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const DEFAULT_SIM_NAME = "KUBESIM"
const DEFAULT_CONFIG_FILE = "/etc/kubedge/kubesim_conf.yaml"
const DEBUG_LOG = false

type Configdata struct {
	Product_name           string   `yaml:"product_name"`
	Product_type           string   `yaml:"product_type"`
	Product_family         string   `yaml:"product_family"`
	Product_release        string   `yaml:"product_release"`
	Feature_set1           []string `yaml:"feature_set1"`
	Feature_set2           []string `yaml:"feature_set2"`
	Enable_log             bool     `yaml:"enable_log"`
	Kubedge_server_port    string   `yaml:"kubedge_server_port"`
	Kubedge_server_address string   `yaml:"kubedge_server_address"`
}

func (config *Configdata) Config(simname string, filename string) {
	if simname == "" {
		simname = DEFAULT_SIM_NAME
	}

	if filename == "" {
		filename = DEFAULT_CONFIG_FILE
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%s: %s: Read: %v", simname, filename, err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("%s: %s: Unmarshal: %v", simname, filename, err)
	}

	if config.Kubedge_server_address == "" {
		config.Kubedge_server_address = "192.168.2.101:30180"
	}

	if config.Kubedge_server_port == "" {
		config.Kubedge_server_port = ":50051"
	}

	if DEBUG_LOG {
		log.Printf("%s: %s contains: product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s, enable_log=%v",
			simname, filename, config.Product_name, config.Product_type, config.Product_family, config.Product_release, config.Feature_set1, config.Feature_set2, config.Enable_log)
	}

}
