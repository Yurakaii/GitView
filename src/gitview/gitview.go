package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	mainWindow := a.NewWindow("GitView")

	mainWindow.Resize(fyne.NewSize(250,550))
	mainWindow.Show()
	a.Run()
}