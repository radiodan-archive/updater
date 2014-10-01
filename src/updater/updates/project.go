package updates

import "updater"

type Project struct {
  Name string
  Refs map[string]updater.Release
}