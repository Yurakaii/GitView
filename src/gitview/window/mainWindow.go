package window

import (
	"gitview/src/gitview/config"
	"gitview/src/gitview/git"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var amend bool
var a fyne.App
var mainWindow fyne.Window
var windowContent *container.Split
const defaultWindowWidth, defaultWindowHeight float32 = 275, 500

func CreateApp() {
	a = app.New()
}

func toolBar() *fyne.Container {
	reloadButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		loadContent() 
	})
	toolbar := container.NewHBox(
		reloadButton,
	)
	return toolbar
}

func topHalf() *fyne.Container {
	changesLabel := widget.NewLabel("Changes:")
	amendCheckBox := widget.NewCheck("Amend", func(b bool) {
		if(b){
			amend = true
		}else{
			amend = false
		}
	})
	content := container.New(
		layout.NewVBoxLayout(),
		toolBar(),
		widget.NewSeparator(),
		changesLabel,
		listChanges(),
		layout.NewSpacer(),
		amendCheckBox,
	)
	return content
}

func bottomHalf() *fyne.Container {
	const placeHolder string = "Commit Message"
	commitMessageEntry := widget.NewEntry()
	commitMessageEntry.SetPlaceHolder(placeHolder)
	forcePushCheckBox := widget.NewCheck("Force Push", func(b bool) {
		if(b){
			config.ForcePush = true
		}else{
			config.ForcePush = false
		}
	})
	commitButton := widget.NewButton("Commit", func() {
		git.Commit(commitMessageEntry.Text, amend)
		loadContent()
	})
	commitPushButton := widget.NewButton("Commit and Push", func() {
		git.CommitAndPush(commitMessageEntry.Text, amend)
		loadContent()
	})
	buttonHBox := container.NewHBox(commitButton, commitPushButton)
	content := container.New(
		layout.NewVBoxLayout(),
		commitMessageEntry,
		layout.NewSpacer(),
		forcePushCheckBox,
		buttonHBox,
	)
	return content
}

func gitContent() *container.Split {
	content := container.NewVSplit(
		topHalf(),
		bottomHalf(),
	)
	return content
}

func loadContent() {
	windowContent = gitContent()
	mainWindow.SetContent(windowContent)
}

func LaunchMainWindow() {
	mainWindow = a.NewWindow("GitView")
	loadContent()
	mainWindow.Resize(fyne.NewSize(defaultWindowWidth, defaultWindowHeight))
	mainWindow.Show()
}

func RunApp() {
	a.Run()
}
