package updates

import "downlo"

type Project struct {
  Name string
  Refs map[string]downlo.Release
}