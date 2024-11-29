package window

import (
	"gitview/src/gitview/git"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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

// TODO: Fix this
func createSmallIcon(resource fyne.Resource) *canvas.Image {
	icon := canvas.NewImageFromResource(resource)
	icon.Resize(fyne.NewSize(16,16))
	return icon
}

func toolBar() *fyne.Container {
	reloadButton := widget.NewButtonWithIcon("", createSmallIcon(theme.ViewRefreshIcon()).Resource, func() {
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
	commitMessageEntry := widget.NewEntry()
	commitButton := widget.NewButton("Commit", func() {git.Commit(commitMessageEntry.Text, amend)})
	commitPushButton := widget.NewButton("Commit and Push", func() {git.CommitAndPush(commitMessageEntry.Text, amend)})
	buttonHBox := container.NewHBox(commitButton, commitPushButton)
	content := container.New(
		layout.NewVBoxLayout(),
		commitMessageEntry,
		layout.NewSpacer(),
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
