package main

import (
	"fyne.io/fyne/v2/app"
	"gitview/src/gitview/window"
)

func main() {
	a := app.New()
	window.LaunchWindow(a)
	a.Run()
}