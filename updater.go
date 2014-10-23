package main

import (
  "flag"
  "fmt"
  "os"
  "os/signal"
  "path/filepath"
  "github.com/radiodan/updater/commands"
  "github.com/radiodan/updater/model"
)

func main() {
  var workspace, target string
  var fs flag.FlagSet

  flag.Parse()

  args := flag.Args()
  command := flag.Arg(0)

  if isValid(command) {
    params := args[1:len(args)]

    fs := flag.NewFlagSet(command, flag.ContinueOnError)
    fs.StringVar(&target, "target", "", "The directory to search for")
    fs.StringVar(&workspace, "workspace", "", "Where to download updates to")

    fs.Parse(params)
  } else {
    fmt.Println("Missing command parameter")
    fmt.Println("Valid commands are: check, install")
    return
  }

  // TODO Make debug flag/env var
  // // debug := false

  // Check arguments
  if target == "" {
    fmt.Println("Missing parameter 'target'")
  }

  if workspace == "" {
    fmt.Println("Missing parameter 'workspace'")
  }

  if target == "" || workspace == "" {
    fs.PrintDefaults()
    return
  }

  // Check or create lock file
  lock, err := model.CreateLockAtPath(filepath.Join(workspace))
  if err != nil {
    fmt.Println("App is already running")
    fmt.Println(err)
    return
  } else {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
      for _ = range c {
        fmt.Println("Release lock")
        lock.Release()
        os.Exit(0)
      }
    }()
    defer lock.Release()
  }

  // Check command
  switch command {
    case "check":
      commands.Check(workspace, target)
    case "install":
      commands.Install(workspace, target)
  }

}

func isValid(command string) (valid bool) {
  valid = false
  if command == "check" || command == "install" {
    valid = true
  }
  return
}