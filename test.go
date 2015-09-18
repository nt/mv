package main

import(
  "io/ioutil"
  "fmt"
  "net/http"
  "log"
  "caffe2"
  "github.com/golang/protobuf/proto"
)


func dumpNet(w http.ResponseWriter, r *http.Request) {
  def := &caffe2.NetDef{}
  data, err := ioutil.ReadFile("inception_net.pb")
  if err != nil {
    log.Fatal(err)
  }
  proto.Unmarshal(data, def)
  proto.MarshalText(w, def)
}

func dumpWeights(w http.ResponseWriter, r *http.Request) {
  def := &caffe2.TensorProtos{}
  data, err := ioutil.ReadFile("inception_tensors.pb")
  if err != nil {
    log.Fatal(err)
  }
  proto.Unmarshal(data, def)
  proto.MarshalText(w, def)
}

func main() {
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", r.URL.Path)
  })
  http.HandleFunc("/net", dumpNet)
  http.HandleFunc("/weights", dumpWeights)

  log.Fatal(http.ListenAndServe(":8080", nil))
}
