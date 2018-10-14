package config
import (
  "log"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)


type Configdata struct {
  Foo string `yaml:"foo"`
  Bar[] string
  Alpha int64 `yaml:"alpha"`
  Beta[] int64 `yaml:"beta"`
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
  //(*config).Foo = "hello world"    //WORKS - change seen in caller
  log.Printf("in library config:  Foo=%s, Bar=%s, Alpha=%d, Beta=%s", config.Foo, config.Bar, config.Alpha, config.Beta)
}
