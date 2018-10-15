package config
import (
  "log"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)


type Configdata struct {
  Product_name string `yaml:"product_name"`
  Product_type string `yaml:"product_type"`
  Product_family string `yaml:"product_family"`
  Product_release string `yaml:"product_release"`
  Feature_set1[] string `yaml:"feature_set1"`
  Feature_set2[] string `yaml:"feature_set2"`
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
  log.Printf("in library config:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
             config.Product_name, config.Product_type, config.Product_family, config.Product_release, config.Feature_set1, config.Feature_set2)

}
