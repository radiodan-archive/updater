package downlo

import (
  "log"
  "strings"
)

func FilterUpdateCandidates(components []Project, latest map[string]map[string]Snapshot) (candidates []Candidate) {
  for _, c := range components {
    if isCandidate, snapshot := NeedsUpdate(c, latest); isCandidate {
      candidate := Candidate{}
      candidate.Url    = snapshot.File
      candidate.Name   = c.Name
      candidate.Ref    = c.Ref
      candidate.Commit = snapshot.Commit
      candidate.Target = c.Path
      candidate.Hash   = snapshot.Sha1
      candidate.FileName = filenameForCandidate(candidate)
      candidates = append(candidates, candidate)
    }
  }
  return
}

func filenameForCandidate(candidate Candidate) string {
  return strings.Replace(candidate.Name, "/", "-", -1) + "-" + candidate.Ref
}

func NeedsUpdate(component Project, latest map[string]map[string]Snapshot) (isCandidate bool, candidate Snapshot) {
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
        candidate = ref
      }
    }
  }

  return
}
