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
	directory := "." // add this to config

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	if(err != nil) {
		panic(err)
	}

	w, err := r.Worktree()
	if(err != nil) {
		panic(err)
	}

	// This is just here for testing, to be removed
	filename := filepath.Join(directory, "example-git-file")
	err = os.WriteFile(filename, []byte("hello world!"), 0644)
	if(err != nil) {
		panic(err)
	}

	// To be removed, staging is going to be done separately
	_, err = w.Add("example-git-file")
	if(err != nil) {
		panic(err)
	}

	// We can verify the current status of the worktree using the method Status.
	status, err := w.Status()
	if(err != nil) {
		panic(err)
	}

	fmt.Println(status)

	// Creates the commit
	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
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
	obj, err := r.CommitObject(commit)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(obj)
}