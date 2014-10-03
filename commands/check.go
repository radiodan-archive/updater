package commands

import (
  "fmt"
  "path/filepath"
  "github.com/radiodan/updater/updates"
  "github.com/radiodan/updater/deployed"
  "github.com/radiodan/updater/model"
)


func Check(workspace string, target string) {

    debug := false

    status := model.StatusFromPath(filepath.Join(workspace, "status.json"))

    fmt.Printf("Scanning %s\n", target)

    deployedReleases := deployed.Deployed(target)
    if debug {
      fmt.Println("deployedReleases")
      fmt.Println(deployedReleases)
    }

    if len(deployedReleases) == 0 {
      fmt.Printf("No deployed releases found in '%s'\n", target)
      return
    }

    latestProjects := updates.LatestReleasesByProject()
    if debug {
      fmt.Println("latestProjects")
      fmt.Println(latestProjects)
    }

    candidates := updates.FilterUpdateCandidates(deployedReleases, latestProjects)
    if debug {
      fmt.Println("candidates")
      fmt.Println(candidates)
    }

    pending := deployed.PendingUpdates(workspace)
    if debug {
      fmt.Println("pending")
      fmt.Println(pending)
    }

    for _, c := range candidates {
      if !updates.Fetched(c, pending, workspace){
        updates.Fetch(c, workspace)
      } else {
        fmt.Printf("%s already downloaded\n", c.Name())
      }
    }

    status.UpdateLastChecked()

    // TODO: Add message/notify
}