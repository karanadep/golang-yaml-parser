package main

import (
	"io/ioutil"
	"path/filepath"
)

func main() {
	filename, _ := filepath.Abs("templates/test.yaml")
	yamlfile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	var config config
}
