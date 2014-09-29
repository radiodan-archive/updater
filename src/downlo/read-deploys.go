package downlo

import (
  "log"
  "path/filepath"
  "io/ioutil"
  "encoding/json"
)

func ReadDeploys(dirs []string) (deploys []Project) {
  for _, dir := range dirs {
    deploys = append(deploys, ReadDeploy(dir))
  }
  return
}

func ReadDeploy(dir string) (deploy Project) {

  path := filepath.Join(dir, ".deploy")

  log.Printf("Reading: %s \n", path)

  contents, err := ioutil.ReadFile(path)
  if err != nil {
    log.Printf("Error reading file: %s \n", path)
  }

  parseError := json.Unmarshal(contents, &deploy)

  deploy.Path = dir

  if parseError != nil {
    log.Printf("Error reading deploy: %s \n", path)
    log.Println(parseError)
  }

  return
}