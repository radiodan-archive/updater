package downlo

import (
  "log"
  "strings"
  "path/filepath"
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func Fetch(candidate Candidate, destinationPath string) () {
  absolutePath, _ := filepath.Abs(destinationPath)
  filename := strings.Replace(candidate.Name, "/", "-", -1) + "-" + candidate.Ref
  downloadPath := filepath.Join(absolutePath, "binaries", filename)

  manifestPath := filepath.Join(absolutePath, "manifests", filename + ".json")

  log.Printf("Download url '%s'\n", candidate.Url)
  log.Printf("Download to '%s'", downloadPath)
  log.Printf("Saving manifest to '%s'", manifestPath)

  candidate.Source = downloadPath

  downloadFile(candidate.Url, downloadPath)
  writeManifest(manifestPath, candidate)

  log.Println(candidate)
}

func downloadFile(url string, path string) {
  resp, err := http.Get(url)
  if err != nil {
    log.Printf("HTTP Request Error '%s'", url)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Printf("Error reading body '%s'", url)
  }
  ioutil.WriteFile(path, body, 0777)
}

func writeManifest(path string, candidate Candidate) {
  data, err := json.MarshalIndent(candidate, "", "  ")

  if err != nil {
    log.Printf("Error making JSON manifest '%o'", candidate)
  }
  err = ioutil.WriteFile(path, data, 0777)
  if err != nil {
    log.Printf("Error writing file '%s'\n", path)
  }
}