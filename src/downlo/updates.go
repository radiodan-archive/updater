package downlo

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type Snapshot struct {
    File    string
    Sha1    string
    Commit  string
    Updated string
}

var data map[string]map[string]Snapshot

func Latest() (data map[string]map[string]Snapshot) {
  url := "http://deploy.radiodan.net"
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("HTTP Request Error")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  data = make(map[string]map[string]Snapshot)

  parseError := json.Unmarshal(body, &data)

  if parseError != nil {
    fmt.Printf("JSON parse error")
  }

  return
}