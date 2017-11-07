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
	if strings.Contains(dir, "typo3conf") {
		temp := strings.Split(dir, "/typo3conf")
		dir = temp[0]
	}

	if pwd.isGitDirPath() {
		return dir
	}
	log.Fatal("No root directory found")
	return ""
}

func main() {
	st := flag.Bool("st", false, "print status")
	br := flag.Bool("br", false, "show branches")
	flag.Parse()

	project := createRepository(setRootDirectory())

	if *br {
		project.getBranches()
	}

	if *st {
		project.getStatus()
	}

	fmt.Println(*project)
}
