package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Status struct {
	filepath string
	data     store
}

type store struct {
	LastChecked time.Time
}

func StatusFromPath(path string) (status Status) {
	var data store

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("No file at '%s' \n", path)
	} else {
		parseError := json.Unmarshal(contents, &data)

		if parseError != nil {
			log.Printf("Error reading: %s \n", path)
			log.Println(parseError)
		}
	}

	status = Status{}
	status.filepath = path
	status.data = data

	return status
}

func (s Status) Save() {
	data, err := json.MarshalIndent(s.data, "", "  ")

	if err != nil {
		log.Printf("Error making JSON '%o'", s.data)
	}
	err = ioutil.WriteFile(s.filepath, data, 0777)
	if err != nil {
		log.Printf("Error writing file '%s'\n", s.filepath)
	}
}

func (s Status) UpdateLastChecked() {
	s.data.LastChecked = time.Now()
	s.Save()
}
