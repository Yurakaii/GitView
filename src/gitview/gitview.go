package main

import (
	"gitview/src/gitview/git"
	"gitview/src/gitview/window"
	
	"fyne.io/fyne/v2/app"
)

func main() {
	git.SetupWorktree()
	a := app.New()
	window.LaunchMainWindow(a)
	a.Run()
}