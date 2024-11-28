package git

import (
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
			Name:  "test", // add this to the config
			Email: "john@doe.org", // add this to the config
			When:  time.Now(),
		},
		Amend: amend,
	})
}

// TODO: Set origin in the config file
func CommitAndPush(name string, amend bool) {
	Commit(name, amend)
	Repo.Push(&git.PushOptions{}) // Default goes to origin, to be changed in the configuration file
}