package downlo

import (
  "log"
  "io/ioutil"
  "os/exec"
)

func InstallUpdate(update Candidate) (bool){

  // create temp dir
  tempDir, err := ioutil.TempDir("", update.FileName + "-")
  if err != nil {
    log.Printf("Error creating temp dir for update '%s'", update.Name)
  }

  log.Printf("Temp Dir %s\n",  tempDir)

  // untar archive
  err = exec.Command("tar", "-C", tempDir, "-xzf", update.Source).Run()
  if err != nil {
    log.Printf("Error unarchiving update '%s'\n", update.Name)
    log.Println(err)
    return false
  }

  // mv old file
  err = exec.Command("mv", update.Target, update.Target + ".previous").Run()
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

    exec.Command("mv", update.Target + ".previous", update.Target).Run()

    return false
  }

  // signal success
  return true
}