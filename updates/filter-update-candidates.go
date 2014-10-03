package updates

import (
  "log"
  "github.com/radiodan/updater/model"
)

func FilterUpdateCandidates(deployed []model.Release, latest []Project) (candidates []model.Release) {
  for _, d := range deployed {
    if isCandidate, release := NeedsUpdate(d, latest); isCandidate {
      r := model.Release{}

      // Take most info from release
      r.Project = release.Project
      r.Ref     = release.Ref
      r.Commit  = release.Commit
      r.Hash    = release.Hash

      // The URL to fetch the release from
      r.Source  = release.Source

      // The app to be deployed
      r.Target  = d.Source

      candidates = append(candidates, r)
    }
  }

  return
}

func NeedsUpdate(deployed model.Release, latest []Project) (isCandidate bool, candidate model.Release) {
  log.Printf("Checking '%s#%s'...\n", deployed.Project, deployed.Ref)

  isCandidate = false

  for _, project := range latest {
    if project.Name == deployed.Project {
      log.Println("...found project")

      for refName, ref := range project.Refs {
        if refName == deployed.Ref {
          log.Printf("...found ref\n")

          if deployed.Commit == ref.Commit {
            log.Printf("...needs updates? NO")
            log.Printf("...(latest: %s, files: %s)\n", deployed.Commit, ref.Commit)
          } else {
            log.Printf("...needs update? YES\n")
            isCandidate = true
            candidate = ref
          }
        }
      }
    }
  }

  return
}
