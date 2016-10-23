package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type repository struct {
	dir           string
	status        string
	currentBranch string
	projectType   string
	branches      []string
	children      []repository
}

func (repo repository) runGitCommand(param ...string) string {
	curDir, _ := os.Getwd()
	os.Chdir(repo.dir)
	defer os.Chdir(curDir)

	cmd := exec.Command("git", param...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	return out.String()
}

func (repo *repository) getStatus() string {
	repo.status = repo.runGitCommand("status", "--porcelain")
	fmt.Println(repo.status)
	return repo.status
}

func (repo *repository) getBranches() []string {
	branches := strings.Split(repo.runGitCommand("branch"), "\n")
	var result []string
	for _, value := range branches {
		isMaster := false
		if strings.Contains(value, "*") {
			isMaster = true
		}
		tempName := strings.Trim(value, " *")
		if len(tempName) > 0 {
			result = append(result, tempName)
		}
		if isMaster {
			repo.currentBranch = tempName
		}
	}
	repo.branches = result

	return repo.branches
}

func (repo *repository) setProjectType() {
	if _, err := os.Stat(repo.dir + "/typo3conf"); err == nil {
		repo.projectType = "TYPO3"
	}
}

func (repo *repository) findChildren() (result []repository) {
	if repo.projectType != "TYPO3" {
		tempDir, _ := os.Getwd()
		os.Chdir(repo.dir + "/typo3conf/ext/")
		defer os.Chdir(tempDir)

	}
	return
}
