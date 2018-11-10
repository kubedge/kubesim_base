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

type Configdata struct {
	Product_name    string   `yaml:"product_name"`
	Product_type    string   `yaml:"product_type"`
	Product_family  string   `yaml:"product_family"`
	Product_release string   `yaml:"product_release"`
	Feature_set1    []string `yaml:"feature_set1"`
	Feature_set2    []string `yaml:"feature_set2"`
	Enable_log      bool     `yaml:"enable_log"`
}

func (config *Configdata) Config() {
	log.Printf("%s", "Config invoked")

	yamlFile, err := ioutil.ReadFile("/etc/kubedge/kubesim_conf.yaml")
	//yamlFile, err := ioutil.ReadFile("/tmp/kubesim_conf.yaml")
	log.Printf("yamlFile=%s", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	//(*config).Product_name = "hello world"    //WORKS - change seen in caller
	log.Printf("in library config:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s, enable_log=%v",
		config.Product_name, config.Product_type, config.Product_family, config.Product_release, config.Feature_set1, config.Feature_set2, config.Enable_log)

}
