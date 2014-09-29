package downlo

type Project struct {
    Name    string
    Ref     string
    Commit  string
    Path    string
}

type Snapshot struct {
    File    string
    Sha1    string
    Commit  string
    Updated string
}

type Candidate struct {
    Url    string
    Name   string
    Ref    string
    Commit string
    Hash   string
    Target string
    Source string
}