package deployed

import (
	"github.com/radiodan/updater/model"
	"log"
	"os/exec"
	"path/filepath"
)

func CleanUp(release model.Release, workspace string) bool {
	var err error

	absolutePath, _ := filepath.Abs(workspace)
	filename := release.Name()
	downloadPath := filepath.Join(absolutePath, "downloads", filename)
	manifestPath := filepath.Join(absolutePath, "manifests", filename+".json")

	err = exec.Command("rm", "-r", downloadPath).Run()
	if err != nil {
		log.Printf("Removing binary '%s'\n", release.Project)
		log.Println(err)
		return false
	}

	err = exec.Command("rm", "-r", manifestPath).Run()
	if err != nil {
		log.Printf("Removing manifest '%s'\n", release.Project)
		log.Println(err)
		return false
	}

	return true
}
