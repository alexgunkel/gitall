package main

import (
	"os"
)

type path struct {
	c string
}

func (p path) add(s string) path {
	return path{p.c + s}
}

func (p path) String() string {
	return p.c
}

type directory struct {
	absolutePath path
}

func (dir directory) getWorkingDir() path {
	var tmp string
	tmp, _ = os.Getwd()
	dir.absolutePath = path{tmp}

	return dir.absolutePath
}

func (dir directory) isGitDirPath() bool {
	gitDir := dir.absolutePath.add("/.git")
	if _, err := os.Stat(gitDir.String()); os.IsNotExist(err) {
		return false
	}

	return true
}

func (dir directory) createRepository() (repo *repository) {
	repo = new(repository)
	repo.dir = dir.absolutePath
	repo.setProjectType()

	return
}
