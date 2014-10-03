package updates

import "github.com/radiodan/updater/model"

type Project struct {
  Name string
  Refs map[string]model.Release
}