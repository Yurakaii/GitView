package git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Used for the commit button, and for the commit and push feature.
// TODO: make a YAML configuration file for git configuration
func Commit(name string) {
	var err error

	// Creates the commit
	commit, err := Worktree.Commit(name, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "test", // add this to the config
			Email: "john@doe.org", // add this to the config
			When:  time.Now(),
		},
	})
	if(err != nil) {
		panic(err)
	}

	// Prints the current HEAD to verify that all worked well.
	obj, err := Repo.CommitObject(commit)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(obj)
}