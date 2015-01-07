package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Lock struct {
	filepath string
}

func (l Lock) Release() {
	os.Remove(l.filepath)
}

func CreateLockAtPath(path string) (lock Lock, err error) {

	absolutePath, _ := filepath.Abs(path)
	pid := fmt.Sprintf("%d\n", os.Getpid())

	lockPath := filepath.Join(absolutePath, "updater.pid")

	_, statErr := os.Stat(lockPath)
	if statErr == nil {
		// File exists
		err = errors.New("Lock file exists " + lockPath)
	} else {
		lock = Lock{}
		lock.filepath = lockPath
		ioutil.WriteFile(lockPath, []byte(pid), 0777)
	}
	return
}
