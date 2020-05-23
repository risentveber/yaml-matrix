package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/risentveber/yaml-matrix/matrix"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: cli <filename>")
		os.Exit(1)
	}

	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println("filename error", err)
		os.Exit(2)
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file error", err)
		os.Exit(3)
	}

	s := matrix.NewService("matrix")
	bytes, err := s.Convert(yamlFile)

	if err != nil {
		fmt.Println("convert error", err)
		os.Exit(4)
	}
	fmt.Println(string(bytes))
}
