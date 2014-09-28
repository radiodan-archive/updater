package downlo

import (
  "fmt"
)

func FilterUpdateCandidates(components []Project, latest map[string]map[string]Snapshot) (candidates []Project) {
  for _, c := range components {
    if NeedsUpdate(c, latest) {
      candidates = append(candidates, c)
    }
  }
  return
}

func NeedsUpdate(component Project, latest map[string]map[string]Snapshot) (isCandidate bool) {
  fmt.Printf("Checking '%s#%s'...\n", component.Name, component.Ref)

  isCandidate = false

  if project, ok := latest[component.Name]; ok {
    fmt.Println("...found project")

    if ref, ok := project[component.Ref]; ok {
      fmt.Printf("...found ref\n")

      if component.Commit == ref.Commit {
        fmt.Printf("...needs updates? NO")
        fmt.Printf("...(latest: %s, files: %s)\n", component.Commit, ref.Commit)
      } else {
        fmt.Printf("...needs update? YES\n")
        isCandidate = true
      }
    }
  }

  return
}
