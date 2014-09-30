package deployed

import (
  "log"
  "io/ioutil"
  "path/filepath"
  "os/exec"
  "downlo"
)

func InstallUpdate(update downlo.Release, workspace string) (bool){
  absolutePath, _ := filepath.Abs(workspace)
  backupPath := filepath.Join(absolutePath, "previous", update.Name)

  // create temp dir
  tempDir, err := ioutil.TempDir("", update.Name)
  if err != nil {
    log.Printf("Error creating temp dir for update '%s'", update.Project)
  }

  log.Printf("Temp Dir %s\n",  tempDir)

  // untar archive
  err = exec.Command("tar", "-C", tempDir, "-xzf", update.Source).Run()
  if err != nil {
    log.Printf("Error unarchiving update '%s'\n", update.Project)
    log.Println(err)
    return false
  }

  // mv old file
  err = exec.Command("mv", update.Target, backupPath).Run()
  if err != nil {
    log.Printf("Error moving current app '%s'\n", update.Target)
    log.Println(err)
    return false
  }

  // mv new into place
  err = exec.Command("mv", tempDir, update.Target).Run()
  if err != nil {
    log.Printf("Error moving updated app into place'%s'\n", update.Target)
    log.Println(err)

    exec.Command("mv", backupPath, update.Target).Run()

    return false
  }

  // signal success
  return true
}