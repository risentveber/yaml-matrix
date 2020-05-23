package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/risentveber/yaml-matrix/matrix"
)

func main() {
	fmt.Println("started")

	filename, _ := filepath.Abs("./tmp/input.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("file error", err)
		return
	}

	s := matrix.NewService("matrix")
	bytes, err := s.Convert(yamlFile)

	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(string(bytes))
}
