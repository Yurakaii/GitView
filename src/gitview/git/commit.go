package git

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Used for the commit button, and for the commit and push feature.
// TODO: make a YAML configuration file for git configuration
func Commit() {
	var err error
	// This is just here for testing, to be removed
	filename := filepath.Join(Directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0644)
	if(err != nil) {
		panic(err)
	}

	// To be removed, staging is going to be done separately
	_, err = Worktree.Add("example-git-file")
	if(err != nil) {
		panic(err)
	}

	// We can verify the current status of the worktree using the method Status.
	status, err := Worktree.Status()
	if(err != nil) {
		panic(err)
	}

	fmt.Println(status)

	// Creates the commit
	commit, err := Worktree.Commit("example go-git commit", &git.CommitOptions{
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