package connected
import (
  "log"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)


type Connecteddata struct {
  Connected[] string `yaml:"connected"`
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
