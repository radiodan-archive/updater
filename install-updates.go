package main

import (
  "flag"
  "log"
  "downlo"
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

    updates := downlo.PendingUpdates(workspace)
    if debug { log.Println(updates) }

    for _, update := range updates {
      log.Printf("Found update '%s'", update.Name)
      if downlo.IsUpdateValid(update.Source, update.Hash) {
        log.Printf("Update '%s' valid", update.Name)
      } else {
        log.Printf("Update '%s' is not valid", update.Name)
      }
    }
}