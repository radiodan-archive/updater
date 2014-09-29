package main

import (
  "flag"
  "fmt"
  "downlo"
)


func main() {
    var target, workspace string

    debug := false

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

    dirs    := downlo.ScanDirs(target)
    if debug { fmt.Println(dirs) }

    deploys := downlo.ReadDeploys(dirs)
    if debug { fmt.Println(deploys) }

    latest  := downlo.Latest()
    if debug { fmt.Println(latest) }

    candidates := downlo.FilterUpdateCandidates(deploys, latest)
    for _, c := range candidates {
      downlo.Fetch(c, workspace)
    }
    if debug {
      fmt.Println(candidates)
    }

    // downloads := downlo.FetchUpdates(candidates)
}