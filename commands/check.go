package commands

import (
  "fmt"
  "github.com/radiodan/updater/updates"
  "github.com/radiodan/updater/deployed"
)


func Check(workspace string, target string) {

    debug := false

    fmt.Printf("Scanning %s\n", target)

    deployedReleases := deployed.Deployed(target)
    if debug {
      fmt.Println("deployedReleases")
      fmt.Println(deployedReleases)
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

    for _, c := range candidates {
      updates.Fetch(c, workspace)
    }

    // TODO: Add message/notify
}