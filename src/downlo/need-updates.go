package downlo

import (
  "log"
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
  log.Printf("Checking '%s#%s'...\n", component.Name, component.Ref)

  isCandidate = false

  if project, ok := latest[component.Name]; ok {
    log.Println("...found project")

    if ref, ok := project[component.Ref]; ok {
      log.Printf("...found ref\n")

      if component.Commit == ref.Commit {
        log.Printf("...needs updates? NO")
        log.Printf("...(latest: %s, files: %s)\n", component.Commit, ref.Commit)
      } else {
        log.Printf("...needs update? YES\n")
        isCandidate = true
      }
    }
  }

  return
}
