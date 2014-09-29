package downlo

import (
  "log"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

var data map[string]map[string]Snapshot

func Latest() (data map[string]map[string]Snapshot) {
  url := "http://deploy.radiodan.net"
  resp, err := http.Get(url)
  if err != nil {
    log.Printf("HTTP Request Error")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  data = make(map[string]map[string]Snapshot)

  parseError := json.Unmarshal(body, &data)

  if parseError != nil {
    log.Printf("JSON parse error")
  }

  return
}