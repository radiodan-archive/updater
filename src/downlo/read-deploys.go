package downlo

import (
  "fmt"
  "path/filepath"
  "io/ioutil"
  "encoding/json"
)

type Project struct {
    Name    string
    Ref     string
    Commit  string
}

func ReadDeploys(dirs []string) (deploys []Project) {
  for _, dir := range dirs {
    deploys = append(deploys, ReadDeploy(dir))
  }
  return
}

func ReadDeploy(dir string) (deploy Project) {

  path := filepath.Join(dir, ".deploy")

  fmt.Printf("Reading: %s \n", path)

  contents, err := ioutil.ReadFile(path)
  if err != nil {
    fmt.Printf("Error reading file: %s \n", path)
  }

  parseError := json.Unmarshal(contents, &deploy)

  if parseError != nil {
    fmt.Printf("Error reading deploy: %s \n", path)
    fmt.Println(parseError)
  }

  return
}