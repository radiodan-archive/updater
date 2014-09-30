package main

import (
  "flag"
  "fmt"
  "downlo/updates"
  "downlo/deployed"
)


func main() {
    var target, workspace string

    debug := true

    flag.StringVar(&target, "target", "", "The directory to search for")
    flag.StringVar(&workspace, "workspace", "", "Where to download updates to")
    flag.Parse()

    if target == "" {
      fmt.Println("Missing parameter 'target'")
      flag.PrintDefaults()
      return
    }

    if workspace == "" {
      fmt.Println("Missing parameter 'workspace'")
      flag.PrintDefaults()
      return
    }

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