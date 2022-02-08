package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func getFiles(aDir string) []string {

	var theFiles []string

	files, err := ioutil.ReadDir("./data/")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		theFiles = append(theFiles, f.Name())

	}

	return theFiles
}

func main() {

	someFiles := getFiles("./data/")

	for _, f := range someFiles {

		fmt.Println(f)

	}

}
