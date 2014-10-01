package main

import (
  "flag"
  "log"
  "updater/deployed"
)

func main() {
    var workspace string

    debug := false

    flag.StringVar(&workspace, "workspace", "", "Where to look for downloaded updates")
    flag.Parse()

    if workspace == "" {
      log.Println("Missing parameter 'workspace'")
      flag.PrintDefaults()
      return
    }

    log.Printf("Scanning '%s' for updates to install\n", workspace)

    pending := deployed.PendingUpdates(workspace)
    if debug { log.Println(pending) }

    if len(pending) == 0 {
      log.Printf("No pending updates\n")
    }

    for _, release := range pending {
      log.Printf("Found update '%s'", release.Project)
      if deployed.IsValidRelease(release) {
        log.Printf("Update '%s' valid", release.Project)
        success := deployed.InstallUpdate(release, workspace)
        if success {
          deployed.CleanUp(release, workspace)
        }
      } else {
        log.Printf("Update '%s' is not valid", release.Project)
      }
    }
}