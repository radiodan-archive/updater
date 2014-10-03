package updates

import (
  "log"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "github.com/radiodan/updater/model"
)

func LatestReleasesByProject() (projects []Project) {

  body, _ := fetch("http://deploy.radiodan.net")

  var data interface{}
  parseError := json.Unmarshal(body, &data)
  if parseError != nil {
    log.Printf("JSON parse error")
  }

  projects = parseJsonToProjects(data)

  return
}

// Fetch body from a URL
func fetch(url string) (body []byte, err error) {
  resp, err := http.Get(url)

  if err != nil {
    log.Printf("HTTP Request Error")
  }
  defer resp.Body.Close()

  body, err = ioutil.ReadAll(resp.Body)
  return
}

func parseJsonToProjects(data interface{}) (projects []Project) {
  parsed := data.(map[string]interface{})

  for projectName, contents := range parsed {

    refs := contents.(map[string]interface{})

    current := Project{}
    current.Name = projectName

    current.Refs = map[string]model.Release{}

    for refName, ref := range refs {
      r := ref.(map[string]interface{})

      release := model.Release{}
      release.Project = projectName
      release.Ref     = refName
      release.Source  = r["file"].(string)
      release.Hash    = r["sha1"].(string)
      release.Commit  = r["commit"].(string)

      current.Refs[refName] = release
    }

    projects = append(projects, current)
  }

  return
}