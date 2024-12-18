package git

import (
	"gitview/src/gitview/config"
	"log"

	"github.com/go-git/go-git/v5"
)

var Directory string = "."
var Repo *git.Repository
var Worktree *git.Worktree

func SetupWorktree() {
	var err error

	// This opens the repository
	Repo, err = git.PlainOpen(config.ReadConfig().Directory)
	if(err != nil) {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// Gets the worktree
	Worktree, err = Repo.Worktree()
	if(err != nil) {
		log.Fatalf("Failed to get worktree: %v", err)
	}
}