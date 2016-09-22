package main

import (
  "fmt"
  "log"
  "os"
  "time"
  "strconv"
  "net/http"
	"math/rand"
	k8s_client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/api"
)

var svc api.Service
func initSvc() {
  svc.Spec.Type = api.ServiceTypeClusterIP
  ports := make([]api.ServicePort, 2)
  ports[0].Name = "echo"
  ports[0].Port = 9998
  ports[1].Name = "helloworld"
  ports[1].Port = 9999
  svc.Spec.Ports = ports
  selector := make(map[string]string)
  selector["k8s-app"] = "echo-server"
  svc.Spec.Selector = selector
}
const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano());
	client, err := k8s_client.NewInCluster()
	if err != nil {
		fmt.Println("Failed to make client:")
		panic(err)
	}
  initSvc()
  svcNum, _ := strconv.Atoi(os.Getenv("SVC_NUMBER"))
  svcNameMap := make(map[string]bool)
  baseName := RandStringBytes(5)
  fmt.Println(baseName)
  for i := 0; i < svcNum; i++ {
    svcName := "echo-" + baseName + strconv.Itoa(i)
    svc.ObjectMeta.Name = svcName
    if _, err := client.Services("default").Create(&svc); err != nil {
      //error when creating this svc
      fmt.Println(err)
      svcNameMap[svcName] = false
      continue
    }
    svcNameMap[svcName] = true
    if i % 10 == 0 {
      time.Sleep(5 * time.Second)
      fmt.Println("num:", i)
      for svcName := range svcNameMap {
        if svcNameMap[svcName] {
          //skip svc not created
          url := "http://" + svcName + ":9998"
          resp, err := http.Get(url)
          if err != nil {
            log.Fatal("error:", err)
            continue
          }
          if resp.StatusCode != 200 {
            fmt.Println("response Status:", resp.Status)
            fmt.Println("response Headers:", resp.Header)
            fmt.Println("response Body:", resp.Body)
          }
        }
      }
    }

  }
  select {

  }
}
