package updates

import (
  "log"
  "path/filepath"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "github.com/radiodan/updater/model"
  "github.com/radiodan/updater/deployed"
)

func Fetch(release model.Release, destinationPath string) () {

  downloadPath, manifestPath := filepaths(destinationPath, release)

  log.Printf("Download url '%s'\n", release.Source)
  log.Printf("Download to '%s'", downloadPath)
  log.Printf("Saving manifest to '%s'", manifestPath)

  downloadFile(release.Source, downloadPath)

  release.Source = downloadPath
  writeManifest(manifestPath, release)

  log.Println(release)
}

func Fetched(release model.Release, pendingReleases []model.Release, destinationPath string) (fetched bool) {
  fetched = false

  for _, pending := range pendingReleases {
    if pending.Name() == release.Name() {
      if pending.Commit == release.Commit {
        if deployed.IsValidRelease(pending) {
          fetched = true
        }
      }
    }
  }

  return
}

func filepaths(destinationPath string, release model.Release) (downloadPath, manifestPath string) {
  absolutePath, _ := filepath.Abs(destinationPath)
  name := release.Name()

  downloadPath = filepath.Join(absolutePath, "downloads", name)
  manifestPath = filepath.Join(absolutePath, "manifests", name + ".json")

  return
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

func writeManifest(path string, release model.Release) {
  data, err := json.MarshalIndent(release, "", "  ")

  if err != nil {
    log.Printf("Error making JSON manifest '%o'", release)
  }
  err = ioutil.WriteFile(path, data, 0777)
  if err != nil {
    log.Printf("Error writing file '%s'\n", path)
  }
}