package model

import "strings"

type Release struct {
	Project string
	Ref     string
	Source  string
	Target  string
	Hash    string
	Commit  string
}

func (r Release) Name() string {
	return strings.Replace(r.Project, "/", "-", -1) + "-" + r.Ref
}
