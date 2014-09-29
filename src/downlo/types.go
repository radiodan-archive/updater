package downlo

type Project struct {
    Name    string
    Ref     string
    Commit  string
}

type Snapshot struct {
    File    string
    Sha1    string
    Commit  string
    Updated string
}