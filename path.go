package main

type directory struct {
	absolutePath string
}

func (path directory) getwd() string {
	path.abs, _ = os.Getwd()

	return path.absolutePath
}

func (path directory) isGitDirPath() bool {
	gitDir := path + "/.git"
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return false
	}

	return true
}

func (path directory) createRepository() (repo *repository) {
	repo = new(repository)
	repo.dir = path
	repo.setProjectType()

	return
}
