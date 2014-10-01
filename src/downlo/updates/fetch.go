package updates

import (
  "log"
  "path/filepath"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "strings"
  "downlo"
)

func Fetch(release downlo.Release, destinationPath string) () {
  absolutePath, _ := filepath.Abs(destinationPath)
  filename := filename(release)
  downloadPath := filepath.Join(absolutePath, "downloads", filename)
  manifestPath := filepath.Join(absolutePath, "manifests", filename + ".json")

  log.Printf("Download url '%s'\n", release.Source)
  log.Printf("Download to '%s'", downloadPath)
  log.Printf("Saving manifest to '%s'", manifestPath)

  downloadFile(release.Source, downloadPath)

  release.Source = downloadPath
  release.Name   = filename
  writeManifest(manifestPath, release)

  log.Println(release)
}

func filename(r downlo.Release) string {
  return strings.Replace(r.Project, "/", "-", -1) + "-" + r.Ref
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

func writeManifest(path string, release downlo.Release) {
  data, err := json.MarshalIndent(release, "", "  ")

  if err != nil {
    log.Printf("Error making JSON manifest '%o'", release)
  }
  err = ioutil.WriteFile(path, data, 0777)
  if err != nil {
    log.Printf("Error writing file '%s'\n", path)
  }
}