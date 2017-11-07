package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func setRootDirectory() string {
	var pwd directory
	dir := pwd.getWorkingDir()

	// check for typo3conf
	if strings.Contains(dir.String(), "typo3conf") {
		temp := strings.Split(dir.String(), "/typo3conf")
		return temp[0]
	}

	log.Fatal("No root directory found")
	return ""
}

func createRepository(rootDir string) *repository  {
	return new(repository)
}

func main() {
	st := flag.Bool("st", false, "print status")
	br := flag.Bool("br", false, "show branches")
	flag.Parse()

	project := createRepository(setRootDirectory())

	if *br {
		fmt.Print(project.dir.String() + ": ")
		fmt.Println(project.getBranches())
	}

	if *st {
		fmt.Print(project.dir.String() + ": ")
		fmt.Println(project.getStatus())
	}
}
