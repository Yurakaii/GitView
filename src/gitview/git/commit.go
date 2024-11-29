package git

import (
	"gitview/src/gitview/config"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Used for the commit button, and for the commit and push feature.
// TODO: make a YAML configuration file for git configuration
func Commit(name string, amend bool) {
	// Creates the commit
	Worktree.Commit(name, &git.CommitOptions{
		Author: &object.Signature{
			Name:  config.ReadConfig().Name,
			Email: config.ReadConfig().Email,
			When:  time.Now(),
		},
		Amend: amend,
	})
}

// TODO: Set origin in the config file
func CommitAndPush(name string, amend bool) {
	Commit(name, amend)
	Repo.Push(&git.PushOptions{
		RemoteURL: config.ReadConfig().RemoteURL,
		Force: config.ReadConfig().ForcePush,
	})
}