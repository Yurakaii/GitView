package git

import (
	"github.com/go-git/go-git/v5"
)

func Stage(filepath string) {
	Worktree.Add(filepath)
}

func UnStage(filepath string) {
	Worktree.Reset(&git.ResetOptions{
		Mode: git.MixedReset,
	})
}