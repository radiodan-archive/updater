package deployed

import (
	"github.com/radiodan/updater/model"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func InstallUpdate(update model.Release, workspace string) bool {
	absolutePath, _ := filepath.Abs(update.Target)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	releasePath := filepath.Join(absolutePath, "releases", timestamp)
	currentPath := filepath.Join(absolutePath, "current")

	// untar archive
	output, err := exec.Command("mkdir", "-p", releasePath).CombinedOutput()
	if err != nil {
		log.Printf("Error creating release dir '%s'", releasePath)
		return false
	}
	log.Printf("Creating release dir '%s'", releasePath)

	log.Printf("Unarchive update '%s' -> '%s'\n", releasePath, update.Source)
	output, err = exec.Command("tar", "-C", releasePath, "-xzf", update.Source).CombinedOutput()
	if err != nil {
		log.Printf("Error unarchiving update '%s' -> '%s' \n", releasePath, update.Source)
		log.Println(err)
		log.Println(string(output))

		// Attempt to delete the deployed app
		exec.Command("rm", "-r", releasePath).Run()
		return false
	}

	// Detect node app and install dependencies if not
	// provided in archive
	log.Println("Checking for package.json")
	packageJsonPath := filepath.Join(releasePath, "package.json")
	if _, err := os.Stat(packageJsonPath); err == nil {
		log.Println("package.json exists")

		nodeModulesPath := filepath.Join(releasePath, "node_modules")
		log.Println("Checking for node_modules")
		if _, err := os.Stat(nodeModulesPath); os.IsNotExist(err) {
			_ = os.Chdir(releasePath)
			log.Println("Run npm install")
			output, err = exec.Command("npm", "install").CombinedOutput()
			if err != nil {
				log.Printf("Error running npm install '%s'\n", releasePath)
				log.Println(err)
				log.Println(string(output[:]))
				return false
			} else {
				log.Println(string(output[:]))
			}
		}
	} else {
		log.Println("package.json does not exist")
	}

	// Remove old symlink
	output, err = exec.Command("rm", currentPath).CombinedOutput()
	if err != nil {
		log.Printf("Error removing current link '%s'\n", currentPath)
		log.Println(err)
		log.Println(output)

		return false
	}

	// symlink new into place
	log.Printf("Symlink '%s %s'\n", releasePath, currentPath)
	output, err = exec.Command("ln", "-s", "-f", releasePath, currentPath).CombinedOutput()
	if err != nil {
		log.Printf("Error moving updated app into place '%s' -> '%s'\n", releasePath, currentPath)
		log.Println(err)

		return false
	}
	log.Println(output)

	// signal success
	return true
}
