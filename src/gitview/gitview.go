package main

import (
	"gitview/src/gitview/git"
	"gitview/src/gitview/window"
)

func main() {
	git.SetupWorktree()
	window.CreateApp()
	window.LaunchMainWindow()
	window.RunApp()
}