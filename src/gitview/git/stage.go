package git

func Stage(filepath string) {
	Worktree.Add(filepath)
}