package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"github.com/alexgunkel/git"
)

func setRootDirectory() string {
	var pwd git.Directory
	dir := pwd.WorkingDir()

	// check for typo3conf
	if strings.Contains(dir.String(), "typo3conf") {
		temp := strings.Split(dir.String(), "/typo3conf")
		return temp[0]
	}

	log.Fatal("No root directory found")
	return ""
}

func createRepository(rootDir string) *git.Repository  {
	return new(git.Repository)
}

func main() {
	st := flag.Bool("st", false, "print status")
	br := flag.Bool("br", false, "show branches")
	flag.Parse()

	project := createRepository(setRootDirectory())

	if *br {
		fmt.Print(project.Dir() + ": ")
		fmt.Println(project.Branches())
	}

	if *st {
		fmt.Print(project.Dir() + ": ")
		fmt.Println(project.Status())
	}
}
