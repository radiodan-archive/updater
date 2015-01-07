package deployed

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func ScanDirs(path string) (dirs []string) {

	dirpath, _ := filepath.Abs(path)

	matches, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Printf("Error reading path")
	}

	for _, f := range matches {
		if f.IsDir() {
			fullFilePath := filepath.Join(dirpath, f.Name())
			dirs = append(dirs, fullFilePath)
		}
	}

	return
}
