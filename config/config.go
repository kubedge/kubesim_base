package config
import (
  "log"
  //"strings"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

//type Configdata struct {
    //Hits int64 `yaml:"hits"`
    //Time int64 `yaml:"time"`
//}

type Configdata struct {
  Foo string
  Bar[] string
}


func (config *Configdata) Config() {
  log.Printf("%s", "Config invoked")

  //yamlFile, err := ioutil.ReadFile("/etc/kubedge/kubesim_conf.yaml")
  //yamlFile, err := ioutil.ReadFile("/tmp/kubesim_conf.yaml")
  yamlFile, err := ioutil.ReadFile("config.yaml")
  log.Printf("yamlFile=%s", string(yamlFile))
  if err != nil {
      log.Printf("yamlFile.Get err   #%v ", err)
  }

  //err = yaml.Unmarshal(yamlFile, &config)
  //err = yaml.Unmarshal(yamlFile, config)
  //err = yaml.Unmarshal(yamlFile, &(*config))
  //err = yaml.Unmarshal(yamlFile, *config)
  err = yaml.Unmarshal(yamlFile, config)
  if err != nil {
      log.Fatalf("Unmarshal: %v", err)
  }
  //(*config).Foo = "hello world"    //WORKS - change seen in caller
  log.Printf("Foo=%s, Bar=%s", config.Foo, config.Bar)
}
