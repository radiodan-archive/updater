package deployed

import (
  "log"
  "encoding/json"
  "path/filepath"
  "io/ioutil"
  "github.com/radiodan/updater/model"
)

func Deployed(path string) (releases []model.Release) {

  dirs := ScanDirs(path)

  for _, d := range dirs {
    release := ReleaseInfoForFilepath(d)
    releases = append(releases, release)
  }

  return
}

func ReleaseInfoForFilepath(dir string) (release model.Release) {

  var data map[string]interface{}

  path := filepath.Join(dir, "current", ".deploy")

  contents, err := ioutil.ReadFile(path)
  if err != nil {
    log.Printf("Error reading file: %s \n", path)
    return
  }

  parseError := json.Unmarshal(contents, &data)

  if parseError != nil {
    log.Printf("Error reading deploy: %s \n", path)
    log.Println(parseError)
    return
  }

  release.Project = data["name"].(string)
  release.Ref     = data["ref"].(string)
  release.Source  = dir
  release.Commit  = data["commit"].(string)

  return
}

func readDeployFile(dir string) (json string) {

  return
}
